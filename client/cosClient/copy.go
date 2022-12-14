package cosClient

import (
	"context"
	"fmt"
)

// copy 需要加上 bucket 前缀
func (c *CosClient) Copy(ctx context.Context, sourceURL, dest string) error {
	sourceURL = fmt.Sprintf("%s.cos.%s.myqcloud.com/%s", c.conf.CosConfig.Bucket, c.conf.CosConfig.Region, sourceURL)
	dest = fmt.Sprintf("%s.cos.%s.myqcloud.com/%s", c.conf.CosConfig.Bucket, c.conf.CosConfig.Region, dest)
	_, _, err := c.Object.Copy(ctx, dest, sourceURL, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *CosClient) Move(ctx context.Context, sourceURL, dest string) error {
	if err := c.Copy(ctx, sourceURL, dest); err != nil {
		return err
	}
	if _, err := c.Object.Delete(ctx, sourceURL, nil); err != nil {
		return err
	}

	return nil
}
