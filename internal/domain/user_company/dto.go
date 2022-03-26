package user_company

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type CreateUserJobExperienceDTO struct {
	UserId           string    `json:"userId"`
	CompanyId        int64     `json:"companyId,omitempty"`
	CompanyName      string    `json:"companyName,omitempty"`
	EmploymentTypeId int       `json:"employmentTypeId"`
	JobTitle         string    `json:"jobTitle"`
	InProgress       int       `json:"inProgress"`
	StartDate        time.Time `json:"startDate"`
	EndDate          time.Time `json:"endDate,omitempty"`
}

func (uj CreateUserJobExperienceDTO) Validate() error {
	return validation.ValidateStruct(&uj,
		validation.Field(&uj.UserId, validation.Required, validation.Length(36, 36)),
		validation.Field(&uj.CompanyId, validation.Min(1)),
		validation.Field(&uj.CompanyName, validation.Length(2, 50)),
		validation.Field(&uj.EmploymentTypeId, validation.Required, validation.Min(1), validation.Max(7)),
		validation.Field(&uj.JobTitle, validation.Required, validation.Length(2, 50)),
		validation.Field(&uj.InProgress, validation.Min(0), validation.Max(1)),
		validation.Field(&uj.StartDate, validation.Required, validation.Max(time.Now())),
		validation.Field(&uj.EndDate, validation.Max(time.Now())),
	)
}
