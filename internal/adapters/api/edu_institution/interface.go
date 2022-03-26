package edu_institution

import (
	"back/internal/domain/education_institution"
	"context"
)

type eduInstitutionService interface {
	GeListByName(ctx context.Context, name string) (*[]education_institution.EducationInstitution, error)
}
