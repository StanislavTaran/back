package company

import (
	_ "back/internal/domain/company"
	"back/internal/httpHelpers/httpResponse"
	"github.com/gin-gonic/gin"
)

const logLocation = "COMPANY CONTROLLER:"

// @Summary Get companies list by name
// @Tags Company
// @Produce      json
// @Param        name   query      string  true  "company name"
// @Success 200 {array} company.Company
// @Failure 400 {object} httpResponse.ResponseError
// @Failure 401
// @Router /company/list [get]
func (h *Handler) getCompaniesListByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Query("name")

		companies, err := h.companyService.GeListByName(ctx, name)
		if err != nil {
			httpResponse.RequestErrCustomMessage(ctx, err, httpResponse.REQ_ERR_USER_NOT_FOUND)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessData(ctx, companies)
	}
}
