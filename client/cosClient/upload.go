package cosClient

import (
	"bytes"
	"context"
	"fmt"
	"go-interface/model"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func (c *CosClient) InitiateMultipartUpload(ctx context.Context, name string) (*model.InitiateMultipartUploadResult, error) {
	m, _, err := c.Object.InitiateMultipartUpload(ctx, name, nil)
	if err != nil {
		return nil, err
	}

	res := &model.InitiateMultipartUploadResult{
		Bucket:   m.Bucket,
		Key:      m.Key,
		UploadId: m.UploadID,
	}
	return res, nil
}

func (c *CosClient) UploadPart(ctx context.Context, uploadID, name, contentMD5 string, partNumber int, data []byte) (string, error) {
	opt := &cos.ObjectUploadPartOptions{}
	if contentMD5 != "" {
		opt.ContentMD5 = contentMD5
	}

	resp, err := c.Object.UploadPart(ctx, name, uploadID, partNumber, bytes.NewReader(data), opt)
	if err != nil {
		return "", err
	}
	PartETag := resp.Header.Get("ETag")

	return PartETag, nil
}

func (c *CosClient) ListParts(ctx context.Context, uploadID, name string) ([]model.UploadList, error) {
	lp, _, err := c.Object.ListParts(ctx, name, uploadID, nil)
	if err != nil {
		return nil, err
	}

	res := make([]model.UploadList, 0)
	for _, part := range lp.Parts {
		res = append(res, model.UploadList{
			PartNum: part.PartNumber,
			Etag:    part.ETag,
		})
	}

	return res, nil
}

func (c *CosClient) CompleteMultipartUpload(ctx context.Context, uploadID, name string, parts []model.UploadList) error {
	objectParts := make([]cos.Object, 0)
	for _, part := range parts {
		objectParts = append(objectParts, cos.Object{
			PartNumber: part.PartNum,
			ETag:       part.Etag,
		})
	}

	opt := &cos.CompleteMultipartUploadOptions{
		Parts: objectParts, // 所有的上传块
	}

	if _, _, err := c.Object.CompleteMultipartUpload(ctx, name, uploadID, opt); err != nil {
		return err
	}

	return nil
}

func (c *CosClient) AbortMultipartUpload(ctx context.Context, uploadID, name string) error {
	if _, err := c.Object.AbortMultipartUpload(ctx, name, uploadID); err != nil {
		return err
	}

	return nil
}

// 高级操作,不管多大的文件都可以上传
func (c *CosClient) Upload(ctx context.Context, name, filepath string) {
	_, _, err := c.Object.Upload(ctx, name, filepath, nil)
	if err != nil {
		return
	}
}

// 简单操作,最大支持5G
func (c *CosClient) Put(ctx context.Context, name, filepath string, contentType, contentMD5 string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: contentType,
			ContentMD5:  contentMD5,
		},
	}

	_, err = c.Object.Put(context.Background(), name, f, opt)
	if err != nil {
		return err
	}

	return nil
}

// 简单操作,最大支持5G
func (c *CosClient) PutFromFile(ctx context.Context, name, filepath string, contentType, contentMD5 string) error {
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: contentType,
			ContentMD5:  contentMD5,
		},
	}

	if _, err := c.Object.PutFromFile(context.Background(), name, filepath, opt); err != nil {
		return err
	}

	return nil
}

func (c *CosClient) GetPresignedURL(ctx context.Context, key, uploadId string, partNum int32) (string, error) {
	query := &url.Values{}
	query.Add("partNumber", fmt.Sprintf("%d", partNum))
	query.Add("uploadId", uploadId)
	header := &http.Header{}
	header.Add("content-type", "contentType")
	opt := &cos.PresignedURLOptions{
		Query:  query,
		Header: header,
	}
	presignedURL, err := c.Object.GetPresignedURL(ctx, http.MethodPut, key, c.conf.CosConfig.SecretId, c.conf.CosConfig.SecretKey, time.Hour, opt)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

func (c *CosClient) PresignedUpload(ctx context.Context, name, filepath string, headerMap map[string]string) {
	// 获取预签名URL
	key := path.Join(filepath, name)
	presignedURL, err := c.Object.GetPresignedURL(ctx, http.MethodPut, key, c.conf.CosConfig.SecretId, c.conf.CosConfig.SecretKey, time.Hour, nil)
	if err != nil {
		return
	}

	// 2. 通过预签名方式上传对象
	fileDate, err := os.ReadFile(key)
	if err != nil {
		return
	}
	req, err := http.NewRequest(http.MethodPut, presignedURL.String(), bytes.NewReader(fileDate))
	if err != nil {
		return
	}

	for k, v := range headerMap {
		req.Header.Set(k, v)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
}

func (c *CosClient) CreateFolder(ctx context.Context, dirPath string) {
	_, err := c.Object.Put(ctx, dirPath, strings.NewReader(""), nil)
	if err != nil {
		return
	}
}

// 查询指定存储桶中正在进行的分块上传信息
func (c *CosClient) ListMultipartUploads(ctx context.Context, prefix string) error {
	opt := &cos.ListMultipartUploadsOptions{
		Prefix: prefix,
	}
	_, _, err := c.Bucket.ListMultipartUploads(ctx, opt)
	if err != nil {
		return err
	}

	return nil
}

func (c *CosClient) IsExist(ctx context.Context, name string) (bool, error) {
	return c.Object.IsExist(ctx, name)
}
