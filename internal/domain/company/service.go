package company

import "context"

type CompanyService struct {
	companyStorage *Storage
}

func NewCompanyService(companyStorage *Storage) *CompanyService {
	return &CompanyService{
		companyStorage: companyStorage,
	}
}

func (c *CompanyService) GeListByName(ctx context.Context, name string) (*[]Company, error) {
	return c.companyStorage.GetListByName(ctx, name)
}
