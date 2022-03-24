package auth

import (
	authDTO "back/internal/domain/auth"
	"back/internal/httpHelpers/httpResponse"
	jwtpackage "back/pkg/jwt"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

const logLocation = "AUTH CONTROLLER:"

// @Summary Sign In
// @Tags Auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      authDTO.Credentials  true  "Sign In"
// @Success 200 {object} jwtpackage.TokenInfo
// @Failure 400 {object} httpResponse.ResponseError
// @Router /auth/signIn [post]
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

// @Summary Refresh Token
// @Tags Auth
// @Accept       json
// @Produce      json
// @Param        refresh token  body      jwtpackage.RT  true  "Refresh Token"
// @Success 200 {object} jwtpackage.TokenInfo
// @Failure 400 {object} httpResponse.ResponseError
// @Failure 401
// @Router /auth/refresh [post]
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

		claims := &jwtpackage.UserClaims{}

		tkn, err := jwt.ParseWithClaims(rt.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtpackage.JWT_SECRET), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
		if !tkn.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims.Destination != jwtpackage.DESTINATION_REFRESH_TOKEN {
			ctx.AbortWithStatus(http.StatusUnauthorized)
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

// @Summary Log Out
// @Tags Auth
// @Accept       json
// @Produce      json
// @Param        refresh token  body      jwtpackage.RT  true  "Log Out"
// @Success 200
// @Failure 400 {object} httpResponse.ResponseError
// @Router /auth/logout [post]
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
