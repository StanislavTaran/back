package user_company

import (
	"back/internal/domain/user_company"
	"context"
)

type userCompanyService interface {
	Create(ctx context.Context, dto user_company.CreateUserJobExperienceDTO) (id int64, err error)
}
