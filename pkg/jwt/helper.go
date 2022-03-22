package jwtpackage

import (
	"back/internal/domain/user"
	cachepackage "back/pkg/cache"
	"back/pkg/logger"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	AUDIENCE_USERS = "users"
	//AUDIENCE_ADMINS = "admins"
)

// TODO retrieve from ENV
const (
	DESTINATION_TOKEN         = "goodzone"
	DESTINATION_REFRESH_TOKEN = "merefa"
)

const (
	TOKEN_EXP_TIME = 60
	JWT_SECRET     = "jkahdlias9gdliuygasodaso7" // TODO retrieve from ENV
)

type UserClaims struct {
	jwt.StandardClaims
	Email       string `json:"email"`
	Destination string `json:"destination"`
}

type TokenInfo struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type RT struct {
	RefreshToken string `json:"refreshToken"`
}

type helper struct {
	Logger  logger.ILogger
	RTCache cachepackage.Repository
}

func NewHelper(RTCache cachepackage.Repository, logger logger.ILogger) Helper {
	return &helper{RTCache: RTCache, Logger: logger}
}

type Helper interface {
	GenerateAccessToken(u user.User) (*TokenInfo, error)
	UpdateRefreshToken(rt RT) (*TokenInfo, error)
	RemoveRefreshTokenFromCache(rt RT) error
}

func (h *helper) UpdateRefreshToken(rt RT) (*TokenInfo, error) {
	defer h.RTCache.Del([]byte(rt.RefreshToken))

	userBytes, err := h.RTCache.Get([]byte(rt.RefreshToken))
	if err != nil {
		return nil, err
	}
	var u user.User
	err = json.Unmarshal(userBytes, &u)
	if err != nil {
		return nil, err
	}
	return h.GenerateAccessToken(u)
}

func (h *helper) GenerateAccessToken(user user.User) (*TokenInfo, error) {
	expirationTime := time.Now().Add(TOKEN_EXP_TIME * time.Minute)
	claims := &UserClaims{
		Email:       user.Email,
		Destination: DESTINATION_TOKEN,
		StandardClaims: jwt.StandardClaims{
			Id:        user.Id,
			ExpiresAt: expirationTime.Unix(),
			Audience:  AUDIENCE_USERS,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return nil, err
	}

	h.Logger.Info("create refresh token")
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		Email:       user.Email,
		Destination: DESTINATION_REFRESH_TOKEN,
		StandardClaims: jwt.StandardClaims{
			Id:        user.Id,
			ExpiresAt: expirationTime.Unix(),
			Audience:  AUDIENCE_USERS,
		},
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return nil, err
	}

	userBytes, _ := json.Marshal(user)
	err = h.RTCache.Set([]byte(refreshTokenString), userBytes, 60*60*12)
	if err != nil {
		h.Logger.Error(err.Error())

		return nil, err
	}
	fmt.Println(h.RTCache.EntryCount())

	tokenInfo := TokenInfo{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
	}

	return &tokenInfo, nil
}

func (h *helper) RemoveRefreshTokenFromCache(rt RT) error {
	if ok := h.RTCache.Del([]byte(rt.RefreshToken)); !ok {
		return errors.New("error when remove refresh token")
	}

	return nil
}
