package user

import (
	userDTO "back/internal/domain/user"
	"back/internal/httpHelpers/httpResponse"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

const logLocation = "USER CONTROLLER:"

func (h *Handler) getUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		user, err := h.userService.FindById(ctx, id)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessData(ctx, user)
	}
}

func (h *Handler) createUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto userDTO.CreateUserDTO
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		err = json.Unmarshal(body, &dto)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		err = dto.Validate()
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		id, err := h.userService.Create(ctx, dto)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessData(ctx, map[string]string{"id": id})
	}
}

func (h *Handler) activateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		err := h.userService.ActivateUser(ctx, id)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessOK(ctx)
	}
}

func (h *Handler) signIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var creds userDTO.Credentials

		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		err = json.Unmarshal(body, &creds)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		token, expTime, err := h.userService.SignIn(ctx, creds)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		http.SetCookie(ctx.Writer, &http.Cookie{
			Name:     "token",
			Value:    token,
			Expires:  expTime,
			HttpOnly: true,
		})
	}
}
