package user

import (
	"context"
)

type storage interface {
	FindById(ctx context.Context, id string) (*User, error)
	CollectUserInfoById(ctx context.Context, id string) (*FullUserInfoOutputDTO, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, dto CreateUserInputDTO) (string, error)

	UpdateAvatar(ctx context.Context, userId, avatarPath string) error
}

type UploadFileInfo struct {
	Location string `json:"location"`
}

type fileStorage interface {
	Upload(ctx context.Context, bucket, fileName, filePath string, contentType string) (*UploadFileInfo, error)
}
