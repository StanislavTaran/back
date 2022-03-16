package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

var EMAIL_REGEXP = regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])")

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUsersDTO struct {
	Query struct {
		Id        string `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		IsActive  bool   `json:"isActive"`
		CreatedAt int    `json:"createdAt"`
		UpdatedAt int    `json:"updatedAt"`
	} `json:"query"`
	Sort struct {
		Id        string `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	}
}

type CreateUserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

func (c CreateUserDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.FirstName, validation.Required, validation.Length(2, 20)),
		validation.Field(&c.LastName, validation.Required, validation.Length(2, 20)),
		validation.Field(&c.Password, validation.Required, validation.Length(10, 50)),
		validation.Field(&c.Email, validation.Required, validation.Match(EMAIL_REGEXP)),
	)
}
