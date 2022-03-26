package user_education

import "time"

type UserEducation struct {
	Id               int       `json:"id"`
	UserId           string    `json:"userId"`
	EduInstitutionId int       `json:"eduInstitutionId"`
	Faculty          string    `json:"faculty"`
	InProgress       int       `json:"inProgress"`
	StartDate        time.Time `json:"startDate"`
	EndDate          time.Time `json:"endDate"`
}
