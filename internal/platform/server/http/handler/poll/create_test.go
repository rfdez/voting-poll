package poll_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/voting-poll/internal/platform/server/http/handler/poll"
	"github.com/rfdez/voting-poll/kit/command/commandmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T) {
	commandBus := new(commandmocks.Bus)
	commandBus.On(
		"Dispatch",
		mock.Anything,
		mock.AnythingOfType("creator.PollCommand"),
	).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.PUT("/polls/:id", poll.CreateHandler(commandBus))

	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		b, err := json.Marshal(map[string]interface{}{
			"title":       "",
			"description": "",
		})
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, "/polls/8a1c5cdc-ba57-445a-994d-aa412d23723f", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		b, err := json.Marshal(map[string]interface{}{
			"title":       "Poll 1",
			"description": "This is a poll",
		})
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, "/polls/8a1c5cdc-ba57-445a-994d-aa412d23723f", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
