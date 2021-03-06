package user_company

import (
	companyDomain "back/internal/domain/company"
	"context"
	"errors"
)

type UserCompanyService struct {
	userCompanyStorage userCompanyStorage
	companyStorage     companyStorage
}

func NewUserCompanyService(userCompanyStorage userCompanyStorage, companyStorage companyStorage) *UserCompanyService {
	return &UserCompanyService{
		userCompanyStorage: userCompanyStorage,
		companyStorage:     companyStorage,
	}
}

func (u *UserCompanyService) Create(ctx context.Context, dto CreateUserJobExperienceInputDTO) (int64, error) {
	if dto.CompanyId == 0 {
		if dto.CompanyName == "" {
			return 0, errors.New("company not provided")
		}
		var companyDto = companyDomain.CreateCompanyInputDTO{
			FullName:    dto.CompanyName,
			ShortName:   "",
			Description: "",
		}

		id, err := u.companyStorage.Create(ctx, companyDto)
		if err != nil {
			return 0, err
		}

		dto.CompanyId = id
	}

	if dto.InProgress == 1 {
		dto.EndDate = dto.StartDate
	}

	id, err := u.userCompanyStorage.Create(ctx, dto)
	if err != nil {
		return 0, err
	}

	return id, nil
}
