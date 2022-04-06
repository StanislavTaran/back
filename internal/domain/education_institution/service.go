package education_institution

import (
	"back/internal/adapters/mysql/education_institution"
	"context"
)

type EducationInstitutionService struct {
	educationInstitutionStorage educationInstitutionStorage
}

func NewEducationInstitutionService(educationInstitutionStorage *education_institution.Storage) *EducationInstitutionService {
	return &EducationInstitutionService{
		educationInstitutionStorage: educationInstitutionStorage,
	}
}

func (c *EducationInstitutionService) GeListByName(ctx context.Context, name string) (*[]EducationInstitution, error) {
	return c.educationInstitutionStorage.GetListByName(ctx, name)
}
