package rabbitmq

import (
	"context"
	"encoding/json"

	voting "github.com/rfdez/voting-poll/internal"
	"github.com/rfdez/voting-poll/kit/event"
	"github.com/wagslane/go-rabbitmq"
)

type EventBus struct {
	publisher *rabbitmq.Publisher
	consumer  rabbitmq.Consumer
}

func NewEventBus(publisher *rabbitmq.Publisher, consumer rabbitmq.Consumer) *EventBus {
	return &EventBus{
		publisher: publisher,
		consumer:  consumer,
	}
}

func (b *EventBus) Publish(ctx context.Context, event event.Event) error {
	return b.publisher.Publish([]byte("hello, world"), []string{"routing_key"})
}

func (b *EventBus) Subscribe(eventType event.Type, handler event.Handler) error {
	return b.consumer.StartConsuming(
		func(d rabbitmq.Delivery) rabbitmq.Action {
			switch event.Type(d.RoutingKey) {
			case voting.VoteDeletedEventType:
				var body map[string]interface{}
				err := json.Unmarshal(d.Body, &body)
				if err != nil {
					return rabbitmq.NackRequeue
				}

				evt := voting.NewVoteDeletedEvent(
					body["poll_id"].(string),
					body["option_id"].(string),
				)

				if err := handler.Handle(context.Background(), evt); err != nil {
					return rabbitmq.NackRequeue
				}

				return rabbitmq.Ack
			default:
				return rabbitmq.Ack
			}
			// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue
		},
		"my_queue",
		[]string{"routing_key", "routing_key_2"},
	)
}
