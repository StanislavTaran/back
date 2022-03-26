package user_company

import (
	"back/internal/domain/company"
	"context"
)

type userCompanyStorage interface {
	Create(ctx context.Context, dto CreateUserJobExperienceDTO) (int64, error)
}

type companyStorage interface {
	Create(ctx context.Context, dto company.CreateCompanyDTO) (id int64, err error)
}
