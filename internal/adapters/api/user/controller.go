package user

import (
	"back/internal/adapters/minio/user"
	userDomain "back/internal/domain/user"
	"back/internal/httpHelpers/httpResponse"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"os"
)

const (
	logLocation = "USER CONTROLLER:"
)

// @Summary Get user by id
// @Tags User
// @Produce      json
// @Param        id   path      string  true  "user id"
// @Success 200 {object} userDomain.User
// @Failure 400 {object} httpResponse.ResponseError
// @Failure 401
// @Router /users/:id [get]
func (h *Handler) getUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		u, err := h.userService.FindById(ctx, id)
		if err != nil {
			httpResponse.RequestErrCustomMessage(ctx, err, httpResponse.REQ_ERR_USER_NOT_FOUND)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessData(ctx, u)
	}
}

// @Summary Get user info by id
// @Description Get user info with info about job experience and education
// @Tags User
// @Produce      json
// @Param        id   path      string  true  "user id"
// @Success 200 {object} userDomain.FullUserInfoDTO
// @Failure 400 {object} httpResponse.ResponseError
// @Failure 401
// @Router /users/:id/profile [get]
func (h *Handler) getUserFullInfoById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		u, err := h.userService.GetFullUserInfoById(ctx, id)
		if err != nil {
			httpResponse.RequestErrCustomMessage(ctx, err, httpResponse.REQ_ERR_USER_NOT_FOUND)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessData(ctx, u)
	}
}

// @Summary Create User
// @Description Create user
// @Tags User
// @Accept       json
// @Produce      json
// @Param        userData  body      userDomain.CreateUserDTO  true  "User data"
// @Success 200 {object} object{id=string}
// @Failure 400 {object} httpResponse.ResponseError
// @Router /users [post]
func (h *Handler) createUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto userDomain.CreateUserDTO
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

// @Summary Upload user avatar
// @Description Upload user avatar
// @Tags User
// @Produce      json
// @Param avatar formData file true "Body avatar"
// @Success 200 {object} userDomain.UploadFileInfo
// @Failure 400 {object} httpResponse.ResponseError
// @Failure 401
// @Router /users/:id/avatar [post]
func (h *Handler) uploadAvatar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, header, err := ctx.Request.FormFile("avatar")
		if err != nil {
			httpResponse.RequestErr(ctx, errors.New(fmt.Sprintf("file err : %s", err.Error())))
			return
		}
		fileSize := header.Size
		if fileSize > user.MAX_ALLOWED_AVATAR_SIZE {
			httpResponse.RequestErr(ctx, errors.New("max allowed filed size is 5MB"))
			return
		}

		filename := header.Filename
		filePath := "tmp/" + filename
		out, err := os.Create(filePath)
		if err != nil {
			httpResponse.InternalErr(ctx, err)
			return
		}
		_, err = io.Copy(out, file)
		if err != nil {
			httpResponse.InternalErr(ctx, err)
			return
		}
		err = out.Close()
		if err != nil {
			h.logger.Warn(logLocation, err)
		}

		defer os.Remove(filePath)

		id, ok := ctx.Get("userId")
		if !ok {
			httpResponse.RequestErr(ctx, errors.New("user id not provided"))
			return
		}

		info, err := h.userService.UploadUserAvatar(
			ctx,
			"users",
			fmt.Sprintf("%s/avatar.jpeg", id.(string)),
			filePath,
			"image/jpeg",
		)
		if err != nil {
			httpResponse.InternalErr(ctx, err)
			return
		}

		httpResponse.SuccessData(ctx, info)
	}
}
