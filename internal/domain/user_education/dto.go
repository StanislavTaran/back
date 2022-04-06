package user_education

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type CreateUserEducationInputDTO struct {
	UserId             string    `json:"userId"`
	EduInstitutionId   int64     `json:"eduInstitutionId,omitempty"`
	EduInstitutionName string    `json:"eduInstitutionName,omitempty"`
	Faculty            string    `json:"faculty"`
	InProgress         int       `json:"inProgress"`
	StartDate          time.Time `json:"startDate"`
	EndDate            time.Time `json:"endDate,omitempty"`
}

func (uj CreateUserEducationInputDTO) Validate() error {
	return validation.ValidateStruct(&uj,
		validation.Field(&uj.UserId, validation.Required, validation.Length(36, 36)),
		validation.Field(&uj.EduInstitutionId, validation.Min(1)),
		validation.Field(&uj.EduInstitutionName, validation.Length(2, 50)),
		validation.Field(&uj.Faculty, validation.Required, validation.Length(2, 50)),
		validation.Field(&uj.InProgress, validation.Min(0), validation.Max(1)),
		validation.Field(&uj.StartDate, validation.Required, validation.Max(time.Now())),
		validation.Field(&uj.EndDate, validation.Max(time.Now())),
	)
}
