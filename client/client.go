package client

import (
	"context"
	"fmt"
	"go-interface/client/cosClient"
	"go-interface/client/obsClient"
	"go-interface/config"
	"go-interface/model"
)

type FileSystem interface {
	InitiateMultipartUpload(ctx context.Context, name string) (*model.InitiateMultipartUploadResult, error)
	GetPresignedURL(ctx context.Context, key, uploadId string, partNum int32) (string, error)
	UploadPart(ctx context.Context, uploadID, name, contentMD5 string, partNumber int, data []byte) (string, error)
	ListParts(ctx context.Context, uploadID, name string) ([]model.UploadList, error)
	CompleteMultipartUpload(ctx context.Context, uploadID, name string, parts []model.UploadList) error
	AbortMultipartUpload(ctx context.Context, uploadID, name string) error
	ListMultipartUploads(ctx context.Context, prefix string) error
	IsExist(ctx context.Context, name string) (bool, error)
	// TODO Move()
	// TODO ReName()
	// TODO Copy()
	// TODO Delete()
	// TODO List()
	Close() error
}

func NewFileSystem(conf config.Config) (FileSystem, error) {
	var fs FileSystem
	var err error
	switch conf.SystemKey {
	case "obs":
		fs, err = obsClient.NewObsClient(conf)
		if err != nil {
			return nil, err
		}
	case "cos":
		fs, err = cosClient.NewCosClient(conf)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("system type is invalid")
	}
	return fs, nil
}
