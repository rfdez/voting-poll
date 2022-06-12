package option

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
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		err := commandBus.Dispatch(ctx, creating.NewOptionCommand(
			ctx.Param("option_id"),
			req.Title,
			req.Description,
			ctx.Param("poll_id"),
		))
		if err != nil {
			switch {
			case errors.IsWrongInput(err):
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			case errors.IsNotFound(err):
				ctx.Status(http.StatusNotFound)
				return
			default:
				ctx.Status(http.StatusInternalServerError)
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
