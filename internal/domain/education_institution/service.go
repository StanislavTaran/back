package education_institution

import (
	"context"
)

type EducationInstitutionService struct {
	educationInstitutionStorage educationInstitutionStorage
}

func NewEducationInstitutionService(educationInstitutionStorage educationInstitutionStorage) *EducationInstitutionService {
	return &EducationInstitutionService{
		educationInstitutionStorage: educationInstitutionStorage,
	}
}

func (c *EducationInstitutionService) GeListByName(ctx context.Context, name string) (*[]EducationInstitution, error) {
	return c.educationInstitutionStorage.GetListByName(ctx, name)
}
