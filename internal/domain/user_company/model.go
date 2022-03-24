package user_company

import "time"

type UserCompany struct {
	Id               int       `json:"id"`
	UserId           string    `json:"userId"`
	CompanyId        int       `json:"companyId"`
	CompanyName      string    `json:"companyName"`
	EmploymentTypeId int       `json:"employmentTypeId"`
	JobTitle         string    `json:"jobTitle"`
	InProgress       int       `json:"inProgress"`
	StartDate        time.Time `json:"startDate"`
	EndDate          time.Time `json:"endDate"`
}
