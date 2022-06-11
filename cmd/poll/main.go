package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	voting "github.com/rfdez/voting-poll/internal"
	"github.com/rfdez/voting-poll/internal/creating"
	"github.com/rfdez/voting-poll/internal/decreasing"
	"github.com/rfdez/voting-poll/internal/deleting"
	"github.com/rfdez/voting-poll/internal/platform/bus/inmemory"
	"github.com/rfdez/voting-poll/internal/platform/bus/rabbitmq"
	"github.com/rfdez/voting-poll/internal/platform/server/http"
	"github.com/rfdez/voting-poll/internal/platform/storage/postgresql"
	mq "github.com/wagslane/go-rabbitmq"
)

func main() {
	var cfg config
	err := envconfig.Process("poll", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgresql.NewConnection(
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
		cfg.DbParams,
	)
	if err != nil {
		log.Fatal(err)
	}

	rabbitMQURI := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", cfg.MqUser, cfg.MqPass, cfg.MqHost, cfg.MqPort, cfg.MqVHost)

	publisher, err := mq.NewPublisher(rabbitMQURI, mq.Config{})
	if err != nil {
		log.Fatal(err)
	}
	defer publisher.Close()

	consumer, err := mq.NewConsumer(rabbitMQURI, mq.Config{})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = rabbitmq.NewEventBus(publisher, consumer)
	)

	var (
		pollRepository   = postgresql.NewPollRepository(db, cfg.DbTimeout)
		optionRepository = postgresql.NewOptionRepository(db, cfg.DbTimeout)
	)

	var (
		creatingService   = creating.NewService(pollRepository, optionRepository)
		decreasingService = decreasing.NewService(pollRepository, optionRepository)
	)

	var (
		createPollCommandHandler   = creating.NewPollCommandHandler(creatingService)
		createOptionCommandHandler = creating.NewOptionCommandHandler(creatingService)
	)

	commandBus.Register(creating.PollCommandType, createPollCommandHandler)
	commandBus.Register(creating.OptionCommandType, createOptionCommandHandler)

	if err := eventBus.Subscribe(
		voting.VoteDeletedEventType,
		deleting.NewDecreaseOptionVotesOnVoteDeleted(decreasingService),
	); err != nil {
		log.Fatal(err)
	}

	if err := eventBus.Subscribe(
		voting.VoteDeletedEventType,
		deleting.NewDecreasePollVotersOnVoteDeleted(decreasingService),
	); err != nil {
		log.Fatal(err)
	}

	ctx, httpSrv := http.NewServer(context.Background(), cfg.HttpHost, cfg.HttpPort, cfg.ShutdownTimeout, commandBus)
	if err := httpSrv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

type config struct {
	// Http Server configuration
	HttpHost        string        `default:""`
	HttpPort        uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser    string        `default:"poll"`
	DbPass    string        `default:"poll"`
	DbHost    string        `default:"localhost"`
	DbPort    uint          `default:"5432"`
	DbName    string        `default:"voting_poll"`
	DbParams  string        `default:""`
	DbTimeout time.Duration `default:"5s"`
	// RabbitMQ configuration
	MqUser  string `default:"poll"`
	MqPass  string `default:"poll"`
	MqHost  string `default:"localhost"`
	MqPort  uint   `default:"5672"`
	MqVHost string `envconfig:"MQVHOST" default:"/"`
}
