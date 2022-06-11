package rabbitmq

import (
	"encoding/json"
	"time"

	voting "github.com/rfdez/voting-poll/internal"
	"github.com/rfdez/voting-poll/internal/errors"
	"github.com/rfdez/voting-poll/kit/event"
)

type data struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	OccurredOn string                 `json:"occurred_on"`
	Attributes map[string]interface{} `json:"attributes"`
}

type evt struct {
	Data data                   `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

func serialize(event event.Event) ([]byte, error) {
	e := evt{
		Data: data{
			ID:         event.ID(),
			Type:       string(event.Type()),
			OccurredOn: event.OccurredOn().Format(time.RFC3339),
			Attributes: event.ToPrimitives(),
		},
		Meta: map[string]interface{}{},
	}

	return json.Marshal(&e)
}

func deserialize(data []byte) (event.Event, error) {
	var d evt
	if err := json.Unmarshal(data, &d); err != nil {
		return nil, err
	}

	eventName := d.Data.Type

	switch event.Type(eventName) {
	case voting.VoteDeletedEventType:
		evt := new(voting.VoteDeletedEvent)
		return evt.FromPrimitives(d.Data.ID, d.Data.Attributes, d.Data.ID, d.Data.OccurredOn)
	case voting.VoteCreatedEventType:
		evt := new(voting.VoteCreatedEvent)
		return evt.FromPrimitives(d.Data.ID, d.Data.Attributes, d.Data.ID, d.Data.OccurredOn)
	default:
		return nil, errors.NewWrongInput("unknown event type")
	}
}
