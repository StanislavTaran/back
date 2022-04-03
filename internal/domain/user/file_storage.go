package user

import (
	"back/pkg/minioClient"
	"context"
)

const (
	MAX_ALLOWED_AVATAR_SIZE = 5 * 1024 * 1024
)

type userFileStorage struct {
	client *minioClient.MinioClient
}

func NewUserFileStorage(client *minioClient.MinioClient) *userFileStorage {
	return &userFileStorage{
		client: client,
	}
}

func (f *userFileStorage) Upload(ctx context.Context, bucket, fileName, filePath string, contentType string) (*UploadFileInfo, error) {
	i, err := f.client.Upload(ctx, bucket, fileName, filePath, contentType)
	if err != nil {
		return nil, err
	}

	return &UploadFileInfo{
		Location: i.Domain + "/" + i.Bucket + "/" + i.Key,
	}, nil
}
