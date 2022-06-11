package rabbitmq

import (
	"context"
	"fmt"
	"reflect"

	"github.com/rfdez/voting-poll/internal/errors"
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
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange("voting-poll"),
	)
}

func (b *EventBus) Subscribe(eventType event.Type, handler event.Handler) error {
	return b.consumer.StartConsuming(
		func(d rabbitmq.Delivery) rabbitmq.Action {
			evt, err := deserialize(d.Body)
			if err != nil {
				if errors.IsWrongInput(err) {
					return rabbitmq.Ack
				}

				return rabbitmq.NackDiscard
			}

			if err := handler.Handle(context.Background(), evt); err != nil {
				if errors.IsWrongInput(err) || errors.IsNotFound(err) {
					return rabbitmq.Ack
				}

				return rabbitmq.NackDiscard
			}

			return rabbitmq.Ack
			// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue
		},
		fmt.Sprintf("voting-app.voting-poll.%s", toSnakeCase(reflect.TypeOf(handler).Name())),
		[]string{string(eventType)},
		rabbitmq.WithConsumeOptionsQueueDurable,
		rabbitmq.WithConsumeOptionsBindingExchangeName(getServiceName(eventType)),
		rabbitmq.WithConsumeOptionsBindingExchangeKind("topic"),
		rabbitmq.WithConsumeOptionsBindingExchangeDurable,
		rabbitmq.WithConsumeOptionsConsumerName(toSnakeCase(reflect.TypeOf(handler).Name())),
	)
}
