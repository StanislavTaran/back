package user

import (
	userDTO "back/internal/domain/user"
	"back/internal/httpHelpers/httpResponse"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

const logLocation = "USER CONTROLLER:"

// @Summary Get user by id
// @Tags User
// @Produce      json
// @Param        id   path      string  true  "user id"
// @Success 200 {object} userDTO.User
// @Failure 400 {object} httpResponse.ResponseError
// @Failure 401
// @Router /users/:id [get]
func (h *Handler) getUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		user, err := h.userService.FindById(ctx, id)
		if err != nil {
			httpResponse.RequestErrCustomMessage(ctx, err, httpResponse.REQ_ERR_USER_NOT_FOUND)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessData(ctx, user)
	}
}

// @Summary Get user info by id
// @Description Get user info with info about job experience and education
// @Tags User
// @Produce      json
// @Param        id   path      string  true  "user id"
// @Success 200 {object} userDTO.FullUserInfoDTO
// @Failure 400 {object} httpResponse.ResponseError
// @Failure 401
// @Router /users/:id/profile [get]
func (h *Handler) getUserFullInfoById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		user, err := h.userService.GetFullUserInfoById(ctx, id)
		if err != nil {
			httpResponse.RequestErrCustomMessage(ctx, err, httpResponse.REQ_ERR_USER_NOT_FOUND)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessData(ctx, user)
	}
}

// @Summary Create User
// @Description Create user
// @Tags User
// @Accept       json
// @Produce      json
// @Param        userData  body      userDTO.CreateUserDTO  true  "User data"
// @Success 200 {object} object{id=string}
// @Failure 400 {object} httpResponse.ResponseError
// @Router /users [post]
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
			httpResponse.RequestErrCustomMessage(ctx, err, httpResponse.VALIDATION_ERR)
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
