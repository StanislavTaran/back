package user

import (
	"back/internal/domain/user"
	"context"
	"time"
)

type userService interface {
	Find(ctx context.Context, query map[string]interface{}) ([]*user.User, error)
	FindById(ctx context.Context, id string) (*user.User, error)
	Create(ctx context.Context, dto user.CreateUserDTO) (id string, err error)
	Update(ctx context.Context, dto user.CreateUserDTO) error
	ActivateUser(ctx context.Context, id string) (err error)
	SignIn(ctx context.Context, credentials user.Credentials) (tokenString string, expirationTime time.Time, err error)
}
