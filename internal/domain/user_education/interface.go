package user_education

import (
	"back/internal/domain/education_institution"
	"context"
)

type userEducationStorage interface {
	Create(ctx context.Context, dto CreateUserEducationDTO) (int64, error)
}

type eduInstitutionStorage interface {
	Create(ctx context.Context, dto education_institution.CreateEducationInstitutionDTO) (id int64, err error)
}
