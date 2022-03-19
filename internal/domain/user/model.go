package user

import (
	"time"
)

type User struct {
	Id          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Email       string    `json:"email"`
	Password    string    `json:"password,omitempty"`
	ShortInfo   string    `json:"shortInfo"`
	RoleId      uint8     `json:"roleId,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
