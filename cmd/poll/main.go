package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/rfdez/voting-poll/internal/creating"
	"github.com/rfdez/voting-poll/internal/platform/bus/inmemory"
	"github.com/rfdez/voting-poll/internal/platform/server/http"
	"github.com/rfdez/voting-poll/internal/platform/storage/postgresql"
)

func main() {
	var cfg config
	err := envconfig.Process("POLL", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	psqlURI := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbParams)
	db, err := sql.Open("postgres", psqlURI)
	if err != nil {
		log.Fatal(err)
	}

	var (
		commandBus = inmemory.NewCommandBus()
	)

	var (
		pollRepository   = postgresql.NewPollRepository(db, cfg.DbTimeout)
		optionRepository = postgresql.NewOptionRepository(db, cfg.DbTimeout)
	)

	var (
		creatingService = creating.NewService(pollRepository, optionRepository)
	)

	var (
		createPollCommandHandler   = creating.NewPollCommandHandler(creatingService)
		createOptionCommandHandler = creating.NewOptionCommandHandler(creatingService)
	)

	commandBus.Register(creating.PollCommandType, createPollCommandHandler)
	commandBus.Register(creating.OptionCommandType, createOptionCommandHandler)

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
	DbParams  string        `default:"sslmode=disable"`
	DbTimeout time.Duration `default:"5s"`
}
