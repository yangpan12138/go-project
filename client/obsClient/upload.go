package obsClient

import (
	"context"
	"fmt"
	"go-interface/model"
	"time"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

func (o *ObsClient) InitiateMultipartUpload(ctx context.Context, name string) (*model.InitiateMultipartUploadResult, error) {
	inputInit := &obs.InitiateMultipartUploadInput{}
	inputInit.Bucket = "youObsBucketName"
	inputInit.Key = name
	m, err := o.ObsClient.InitiateMultipartUpload(inputInit)
	if err != nil {
		return nil, err
	}

	res := &model.InitiateMultipartUploadResult{
		Bucket:    m.Bucket,
		Key:       m.Key,
		UploadId:  m.UploadId,
		RequestId: m.RequestId,
	}
	return res, nil
}

func (o *ObsClient) GetPresignedURL(ctx context.Context, key, uploadId string, partNum int32) (string, error) {
	queryParams := make(map[string]string)
	queryParams["partNumber"] = fmt.Sprintf("%d", partNum)
	queryParams["uploadId"] = uploadId

	header := make(map[string]string)
	header["x-obs-date"] = obs.FormatUtcToRfc1123(time.Now().UTC())
	header["content-type"] = "contentType"
	// TODO 列出上传、下载、删除不同的请求 URL
	// 生成上传对象的带授权信息的URL
	putObjectInput := &obs.CreateSignedUrlInput{}
	putObjectInput.Method = obs.HttpMethodPut
	putObjectInput.Bucket = o.conf.ObsConfig.Bucket
	putObjectInput.Key = key
	putObjectInput.Expires = 3600
	putObjectInput.QueryParams = queryParams
	putObjectInput.Headers = header
	putObjectOutput, err := o.ObsClient.CreateSignedUrl(putObjectInput)
	if err == nil {
		return "", err
	}
	fmt.Printf("SignedUrl:%s\n", putObjectOutput.SignedUrl)
	fmt.Printf("ActualSignedRequestHeaders:%v\n", putObjectOutput.ActualSignedRequestHeaders)

	// 生成下载对象的带授权信息的URL
	getObjectInput := &obs.CreateSignedUrlInput{}
	getObjectInput.Method = obs.HttpMethodGet
	getObjectInput.Bucket = o.conf.ObsConfig.Bucket
	getObjectInput.Key = key
	getObjectInput.Expires = 3600
	getObjectInput.QueryParams = queryParams
	getObjectInput.Headers = header
	getObjectOutput, err := o.ObsClient.CreateSignedUrl(getObjectInput)
	if err == nil {
		return "", err
	}
	fmt.Printf("SignedUrl:%s\n", getObjectOutput.SignedUrl)
	fmt.Printf("ActualSignedRequestHeaders:%v\n", getObjectOutput.ActualSignedRequestHeaders)

	// 生成删除对象的带授权信息的URL
	deleteObjectInput := &obs.CreateSignedUrlInput{}
	deleteObjectInput.Method = obs.HttpMethodDelete
	deleteObjectInput.Bucket = o.conf.ObsConfig.Bucket
	deleteObjectInput.Key = key
	deleteObjectInput.Expires = 3600
	deleteObjectInput.QueryParams = queryParams
	deleteObjectInput.Headers = header
	deleteObjectOutput, err := o.ObsClient.CreateSignedUrl(deleteObjectInput)
	if err != nil {
		return "", err
	}
	fmt.Printf("SignedUrl:%s\n", deleteObjectOutput.SignedUrl)
	fmt.Printf("ActualSignedRequestHeaders:%v\n", deleteObjectOutput.ActualSignedRequestHeaders)
	return "", nil
}

func (o *ObsClient) UploadPart(ctx context.Context, uploadID, name, contentMD5 string, partNumber int, data []byte) (string, error) {

	return "", nil
}

func (o *ObsClient) ListParts(ctx context.Context, uploadID, name string) ([]model.UploadList, error) {

	return nil, nil
}

func (o *ObsClient) CompleteMultipartUpload(ctx context.Context, uploadID, name string, parts []model.UploadList) error {

	return nil
}

func (o *ObsClient) AbortMultipartUpload(ctx context.Context, uploadID, name string) error {

	return nil
}

func (o *ObsClient) Upload(ctx context.Context, name, filepath string) {

}

func (o *ObsClient) ListMultipartUploads(ctx context.Context, prefix string) error {

	return nil
}

func (o *ObsClient) IsExist(ctx context.Context, name string) (bool, error) {
	return true, nil
}
