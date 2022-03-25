package company

import (
	"back/internal/domain/company"
	"context"
)

type companyService interface {
	GeListByName(ctx context.Context, name string) (*[]company.Company, error)
}
