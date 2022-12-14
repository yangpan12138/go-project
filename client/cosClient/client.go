package cosClient

import (
	"fmt"
	"go-interface/config"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
)

type CosClient struct {
	*cos.Client
	conf config.Config
}

func NewCosClient(conf config.Config) (*CosClient, error) {
	bucketURL := fmt.Sprintf("http://%s.cos.%s.myqcloud.com", conf.CosConfig.Bucket, conf.CosConfig.Region)
	u, err := url.Parse(bucketURL)
	if err != nil {
		return nil, err
	}

	secretId := "secretId"
	secretKey := "secretKey"

	b := &cos.BaseURL{BucketURL: u}
	transport := &debug.DebugRequestTransport{
		RequestHeader:  true,
		RequestBody:    true,
		ResponseHeader: true,
		ResponseBody:   true,
	}

	if !conf.CosConfig.IsDebug {
		transport = nil
	}

	cosClient := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretId,
			SecretKey: secretKey,
			Expire:    1 * time.Hour,
			Transport: transport,
		},
	})

	return &CosClient{cosClient, conf}, nil
}

func (c *CosClient) Close() error {
	return nil
}
