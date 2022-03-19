package user

import (
	"context"
)

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

func (us *Service) GetFullUserInfoById(ctx context.Context, id string) (*FullUserInfoDTO, error) {
	user, err := us.userStorage.CollectUserInfoById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *Service) Create(ctx context.Context, dto CreateUserDTO) (id string, err error) {
	return us.userStorage.Create(ctx, dto)
}

func (us *Service) Update(ctx context.Context, dto CreateUserDTO) error {
	return nil
}

func (us *Service) ActivateUser(ctx context.Context, id string) (err error) {
	return us.userStorage.ActivateUser(ctx, id)
}
