package user

import (
	"time"
)

type User struct {
	Id          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	CountryCode int       `json:"countryCode"`
	RegionCode  int       `json:"regionCode"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
