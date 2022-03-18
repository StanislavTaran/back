package auth

import (
	authDTO "back/internal/domain/auth"
	"back/internal/httpHelpers/httpResponse"
	jwtpackage "back/pkg/jwt"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

const logLocation = "AUTH CONTROLLER:"

func (h *Handler) signIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var creds authDTO.Credentials

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

		user, err := h.authService.FindUserByEmail(ctx, creds.Email)
		if err != nil {
			httpResponse.RequestErrCustomMessage(ctx, err, httpResponse.REQ_ERR_USER_NOT_FOUND)
			h.logger.Error(logLocation + err.Error())
			return
		}

		token, err := h.authService.SignIn(ctx, *user, creds)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessData(ctx, &token)
	}
}

func (h *Handler) refreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rt jwtpackage.RT

		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		err = json.Unmarshal(body, &rt)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		token, err := h.authService.RefreshToken(ctx, rt)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessData(ctx, &token)
	}
}

func (h *Handler) logOut() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rt jwtpackage.RT

		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		err = json.Unmarshal(body, &rt)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		err = h.authService.LogOut(ctx, rt)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		ctx.Status(http.StatusOK)
	}
}
