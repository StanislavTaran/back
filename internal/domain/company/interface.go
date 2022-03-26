package company

import (
	"context"
)

type companyStorage interface {
	Create(ctx context.Context, dto CreateCompanyDTO) (id int64, err error)
	GetListByName(ctx context.Context, name string) (*[]Company, error)
}
