package user_company

import (
	"back/internal/domain/company"
	"context"
)

type userCompanyStorage interface {
	Create(ctx context.Context, dto CreateUserJobExperienceInputDTO) (int64, error)
}

type companyStorage interface {
	Create(ctx context.Context, dto company.CreateCompanyInputDTO) (id int64, err error)
}
