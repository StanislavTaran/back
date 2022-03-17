package auth

import (
	"back/internal/domain/user"
	jwtpackage "back/pkg/jwt"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userStorage *user.Storage
	jwtHelper   jwtpackage.Helper
}

func NewAuthService(storage *user.Storage, helper jwtpackage.Helper) *Service {
	return &Service{
		userStorage: storage,
		jwtHelper:   helper,
	}
}

func (as *Service) SignIn(ctx context.Context, user user.User, credentials Credentials) (tokenInfo *jwtpackage.TokenInfo, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return nil, err
	}

	token, err := as.jwtHelper.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (as *Service) FindUserByEmail(ctx context.Context, email string) (user *user.User, err error) {
	user, err = as.userStorage.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (as *Service) RefreshToken(ctx context.Context, rt jwtpackage.RT) (tokenInfo *jwtpackage.TokenInfo, err error) {
	tokenInfo, err = as.jwtHelper.UpdateRefreshToken(rt)
	if err != nil {
		return nil, err
	}

	return tokenInfo, nil
}
