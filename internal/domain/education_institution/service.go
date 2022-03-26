package education_institution

import "context"

type EducationInstitutionService struct {
	educationInstitutionStorage *Storage
}

func NewEducationInstitutionService(educationInstitutionStorage *Storage) *EducationInstitutionService {
	return &EducationInstitutionService{
		educationInstitutionStorage: educationInstitutionStorage,
	}
}

func (c *EducationInstitutionService) GeListByName(ctx context.Context, name string) (*[]EducationInstitution, error) {
	return c.educationInstitutionStorage.GetListByName(ctx, name)
}
