package edu_institution

import (
	_ "back/internal/domain/education_institution"
	"back/internal/httpHelpers/httpResponse"
	"github.com/gin-gonic/gin"
)

const logLocation = "EDUCATION INSTITUTION CONTROLLER:"

// @Summary Get Education Institutions list by name
// @Tags Education Institution
// @Produce      json
// @Param        name   query      string  true  "institution name"
// @Success 200 {array} education_institution.EducationInstitution
// @Failure 400 {object} httpResponse.ResponseError
// @Failure 401
// @Router /edu-inst/list [get]
func (h *Handler) getEduInstitutionsListByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Query("name")

		companies, err := h.eduInstitutionService.GeListByName(ctx, name)
		if err != nil {
			httpResponse.RequestErrCustomMessage(ctx, err, httpResponse.REQ_ERR_USER_NOT_FOUND)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessData(ctx, companies)
	}
}
