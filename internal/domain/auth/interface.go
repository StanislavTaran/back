package auth

import (
	"back/internal/domain/user"
	"context"
)

type userStorage interface {
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
}
