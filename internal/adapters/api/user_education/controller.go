package user_education

import (
	"back/internal/domain/user_education"
	"back/internal/httpHelpers/httpResponse"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

const logLocation = "USER EDUCATION CONTROLLER:"

// @Summary Create User Education Experience
// @Description Create User Education Experience
// @Tags User Education
// @Accept       json
// @Produce      json
// @Param        userData  body      user_education.CreateUserEducationDTO  true  "User education data"
// @Success 200 {object} object{id=number}
// @Failure 400 {object} httpResponse.ResponseError
// @Failure 401
// @Router /userEducation [post]
func (h *Handler) createUserEducationExperience() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto user_education.CreateUserEducationDTO
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

		userId, ok := ctx.Get("userId")
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// TODO Move this part to middleware
		if dto.UserId != userId {
			httpResponse.RequestErrCustomMessage(ctx, errors.New("user id in request is not the same with user id"), "operation not allowed")
			return
		}

		id, err := h.userEducationService.Create(ctx, dto)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessData(ctx, map[string]int64{"id": id})
	}
}
