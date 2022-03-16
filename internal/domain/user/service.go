package user

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	TOKEN_EXP_TIME = 60
)

var jwtKey = []byte("8ashdoasdn8asdv3")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Service struct {
	userStorage *Storage
}

func NewUserService(storage *Storage) *Service {
	return &Service{
		userStorage: storage,
	}
}

func (us *Service) Find(ctx context.Context, query map[string]interface{}) ([]*User, error) {
	return nil, nil
}

func (us *Service) FindById(ctx context.Context, id string) (*User, error) {
	user, err := us.userStorage.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *Service) Create(ctx context.Context, dto CreateUserDTO) (id string, err error) {
	return us.userStorage.Create(ctx, dto)
}

func (us *Service) SignIn(ctx context.Context, credentials Credentials) (tokenString string, expirationTime time.Time, err error) {
	userPassWord, err := us.userStorage.GetUserPassword(ctx, credentials.Email)
	var timeNow = time.Now()
	if err != nil {
		return "", timeNow, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userPassWord), []byte(credentials.Password))
	if err != nil {
		return "", timeNow, err
	}

	expirationTime = timeNow.Add(TOKEN_EXP_TIME * time.Minute)

	claims := &Claims{
		Username: credentials.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	if err != nil {
		return "", timeNow, err
	}

	return tokenString, expirationTime, nil
}

func (us *Service) Update(ctx context.Context, dto CreateUserDTO) error {
	return nil
}

func (us *Service) ActivateUser(ctx context.Context, id string) (err error) {
	return us.userStorage.ActivateUser(ctx, id)
}
