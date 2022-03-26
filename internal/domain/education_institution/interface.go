package education_institution

import (
	"context"
)

type educationInstitutionStorage interface {
	Create(ctx context.Context, dto CreateEducationInstitutionDTO) (id int64, err error)
	GetListByName(ctx context.Context, name string) (*[]EducationInstitution, error)
}
