package user

import "context"

type storage interface {
	FindById(ctx context.Context, id string) (*User, error)
	CollectUserInfoById(ctx context.Context, id string) (*FullUserInfoDTO, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, dto CreateUserDTO) (string, error)
}
