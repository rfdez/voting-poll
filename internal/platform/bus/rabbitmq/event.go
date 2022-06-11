package rabbitmq

import (
	"context"
	"fmt"
	"reflect"

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

func (b *EventBus) Publish(_ context.Context, event event.Event) error {
	body, err := serialize(event)
	if err != nil {
		return err
	}

	return b.publisher.Publish(
		body,
		[]string{string(event.Type())},
		rabbitmq.WithPublishOptionsExchange("voting-poll"),
	)
}

func (b *EventBus) Subscribe(eventType event.Type, handler event.Handler) error {
	return b.consumer.StartConsuming(
		func(d rabbitmq.Delivery) rabbitmq.Action {
			evt, err := deserialize(d.Body)
			if err != nil {
				return rabbitmq.NackRequeue
			}

			if err := handler.Handle(context.Background(), evt); err != nil {
				return rabbitmq.NackRequeue
			}

			return rabbitmq.Ack
			// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue
		},
		fmt.Sprintf("voting-app.voting-poll.%s", toSnakeCase(reflect.TypeOf(handler).Name())),
		[]string{string(eventType)},
	)
}
