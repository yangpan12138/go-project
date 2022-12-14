package cosClient

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// 删除文件删除文件夹统一方法
func (c *CosClient) Delete(ctx context.Context, name string) error {
	if _, err := c.Object.Delete(ctx, name, nil); err != nil {
		return err
	}
	return nil
}

func (c *CosClient) DeleteMulti(ctx context.Context, names []string) error {
	obs := []cos.Object{}
	for _, name := range names {
		obs = append(obs, cos.Object{Key: name})
	}

	opt := &cos.ObjectDeleteMultiOptions{
		Objects: obs,
	}

	d, _, err := c.Object.DeleteMulti(ctx, opt)
	if err != nil {
		return err
	}

	fmt.Println("errors", d.Errors)
	return nil
}

func (c *CosClient) DeleteFolderAndObject(ctx context.Context, dirPath string) error {
	var marker string
	opt := &cos.BucketGetOptions{
		Prefix:  dirPath,
		MaxKeys: 1000,
	}

	isTruncated := true
	for isTruncated {
		opt.Marker = marker
		v, _, err := c.Bucket.Get(ctx, opt)
		if err != nil {
			return err
		}
		for _, content := range v.Contents {
			_, err = c.Object.Delete(ctx, content.Key)
			if err != nil {
				log.Error().Err(err).Str("key", content.Key).Msg("删除失败")
				continue
			}
		}
		isTruncated = v.IsTruncated
		marker = v.NextMarker
	}

	return nil
}
