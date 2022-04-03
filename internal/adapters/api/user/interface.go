package user

import (
	"back/internal/domain/user"
	"context"
)

type userService interface {
	FindById(ctx context.Context, id string) (*user.User, error)
	GetFullUserInfoById(ctx context.Context, id string) (*user.FullUserInfoDTO, error)
	Create(ctx context.Context, dto user.CreateUserDTO) (id string, err error)

	UploadUserAvatar(ctx context.Context, bucketName, fileName, filePath string, contentType string) (*user.UploadFileInfo, error)
}
