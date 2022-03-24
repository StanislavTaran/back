package user

import (
	"github.com/guregu/null"
	"time"
)

type User struct {
	Id          string      `json:"id"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	DateOfBirth null.Time   `json:"dateOfBirth" swaggertype:"primitive,string"`
	Email       string      `json:"email"`
	Password    string      `json:"password,omitempty"`
	ShortInfo   null.String `json:"shortInfo" swaggertype:"primitive,string"`
	RoleId      uint8       `json:"roleId,omitempty"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}
