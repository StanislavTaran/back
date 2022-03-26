package user

import (
	"back/internal/domain/user"
	"context"
)

type userService interface {
	Find(ctx context.Context, query map[string]interface{}) ([]*user.User, error)
	FindById(ctx context.Context, id string) (*user.User, error)
	GetFullUserInfoById(ctx context.Context, id string) (*user.FullUserInfoDTO, error)
	Create(ctx context.Context, dto user.CreateUserDTO) (id string, err error)
	Update(ctx context.Context, dto user.CreateUserDTO) error
}
