package user

import (
	"context"
)

type Service struct {
	userStorage storage
	fileStorage fileStorage
}

func NewUserService(storage storage, fileStorage fileStorage) *Service {
	return &Service{
		userStorage: storage,
		fileStorage: fileStorage,
	}
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

func (us *Service) UploadUserAvatar(ctx context.Context, bucketName, fileName, filePath string, contentType string) (*UploadFileInfo, error) {
	info, err := us.fileStorage.Upload(ctx, bucketName, fileName, filePath, contentType)
	if err != nil {
		return nil, err
	}

	userId := ctx.Value("userId")

	err = us.userStorage.UpdateAvatar(ctx, userId.(string), info.Location)
	if err != nil {
		return nil, err
	}

	return info, nil
}
