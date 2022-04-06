package user_education

import (
	"back/internal/domain/user_education"
	"context"
)

type userEducationService interface {
	Create(ctx context.Context, dto user_education.CreateUserEducationInputDTO) (id int64, err error)
}
