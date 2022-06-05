package rabbitmq_test

import (
	"testing"

	"github.com/NeowayLabs/wabbit/amqptest/server"
	"github.com/stretchr/testify/require"
	"github.com/wagslane/go-rabbitmq"
)

func Test_EventBus_Publish_Succeed(t *testing.T) {
	fakeServer := server.NewServer("amqp://localhost:5672/%2f")
	err := fakeServer.Start()
	require.NoError(t, err)
	defer fakeServer.Stop()

	publisher, err := rabbitmq.NewPublisher("amqp://localhost:5672/%2f", rabbitmq.Config{})
	require.NoError(t, err)
	defer publisher.Close()
}
