package poll

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/voting-poll/internal/creating"
	"github.com/rfdez/voting-poll/internal/errors"
	"github.com/rfdez/voting-poll/kit/command"
)

type createRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

// CreateHandler returns an HTTP handler to perform health checks.
func CreateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := commandBus.Dispatch(ctx, creating.NewPollCommand(
			ctx.Param("poll_id"),
			req.Title,
			req.Description,
		))
		if err != nil {
			switch {
			case errors.IsWrongInput(err):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
