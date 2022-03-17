package auth

import (
	"back/internal/domain/auth"
	"back/internal/domain/user"
	jwtpackage "back/pkg/jwt"
	"context"
)

type authService interface {
	SignIn(ctx context.Context, user user.User, creds auth.Credentials) (tokenInfo *jwtpackage.TokenInfo, err error)
	LogOut(ctx context.Context, rt jwtpackage.RT) (err error)
	RefreshToken(ctx context.Context, rt jwtpackage.RT) (tokenInfo *jwtpackage.TokenInfo, err error)
	FindUserByEmail(ctx context.Context, email string) (user *user.User, err error)
}
