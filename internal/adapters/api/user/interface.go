package user

import (
	"back/internal/domain/user"
	"context"
)

type userService interface {
	Find(ctx context.Context, query map[string]interface{}) ([]*user.User, error)
	FindById(ctx context.Context, id string) (*user.User, error)
	Create(ctx context.Context, dto user.CreateUserDTO) (*user.User, error)
	Update(ctx context.Context, dto user.CreateUserDTO) (*user.User, error)
}
