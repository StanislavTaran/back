package user

import (
	"back/internal/httpHelpers/httpResponse"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		user, err := h.userService.FindById(ctx, id)

		if err != nil {
			httpResponse.InternalErr(ctx, err)
			return
		}

		httpResponse.SuccessData(ctx, user)
	}

}
