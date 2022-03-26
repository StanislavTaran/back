package company

import "context"

type CompanyService struct {
	companyStorage companyStorage
}

func NewCompanyService(companyStorage companyStorage) *CompanyService {
	return &CompanyService{
		companyStorage: companyStorage,
	}
}

func (c *CompanyService) GeListByName(ctx context.Context, name string) (*[]Company, error) {
	return c.companyStorage.GetListByName(ctx, name)
}
