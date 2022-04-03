package minioClient

import (
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Config struct {
	ConnectionURL string `json:"connectionUrl"`
	AccessKeyId   string `json:"accessKeyId"`
	SecretKey     string `json:"secretKey"`
	UseSSL        bool   `json:"useSSL"`
	MaxRetries    int    `json:"maxRetries"`
	WaitRetry     int    `json:"waitRetry"`
}

type MinioClient struct {
	Config  *Config
	Storage *minio.Client
}

type UploadFileInfo struct {
	minio.UploadInfo
	Domain string `json:"domain"`
}

func NewMinioClient(cfg *Config) *MinioClient {
	return &MinioClient{
		Config: cfg,
	}
}

func (m *MinioClient) Connect() error {
	client, err := minio.New(m.Config.ConnectionURL, &minio.Options{
		Creds:  credentials.NewStaticV4(m.Config.AccessKeyId, m.Config.SecretKey, ""),
		Secure: m.Config.UseSSL,
	})
	if err != nil {
		return err
	}

	m.Storage = client
	return nil
}

func (m *MinioClient) MakeBucket(ctx context.Context, bucketName string) error {
	err := m.Storage.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := m.Storage.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			return errors.New(fmt.Sprintf("We already own %s\n", bucketName))
		} else {
			return err
		}
	}
	return nil
}

func (m *MinioClient) Upload(ctx context.Context, bucketName, fileName, filePath string, contentType string) (*UploadFileInfo, error) {

	info, err := m.Storage.FPutObject(ctx, bucketName, fileName, filePath, minio.PutObjectOptions{ContentType: contentType})

	// TODO MAKE BUCKET IF NOT EXISTS
	if err != nil {
		exists, _ := m.Storage.BucketExists(ctx, bucketName)
		if !exists {
			err = m.MakeBucket(ctx, bucketName)
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	}
	return &UploadFileInfo{
		info,
		m.Config.ConnectionURL,
	}, nil
}
