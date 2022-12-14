package cosClient

import (
	"context"
	"fmt"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type FileInfo struct {
	Key          string `json:",omitempty"`
	Type         int    // 0 文件夹 1 文件
	ETag         string `json:",omitempty"`
	Size         int64  `json:",omitempty"`
	PartNumber   int    `json:",omitempty"`
	LastModified string `json:",omitempty"`
	StorageClass string `json:",omitempty"`
	VersionId    string `json:",omitempty"`
}

func (c *CosClient) Get(ctx context.Context, prefix string, showAll bool) ([]FileInfo, error) {
	var marker string
	deliter := "/"
	if showAll {
		deliter = ""
	}
	opt := &cos.BucketGetOptions{
		Prefix:    prefix,  // prefix表示要查询的文件夹,不传该字段，表示查询所有
		Delimiter: deliter, // deliter表示分隔符, 设置为/表示列出当前目录下的object, 设置为空表示列出所有的object
		MaxKeys:   1000,    // 设置最大遍历出多少个对象, 一次listobject最大支持1000
	}
	list := make([]FileInfo, 0)
	isTruncated := true
	for isTruncated {
		opt.Marker = marker
		v, _, err := c.Bucket.Get(ctx, opt)
		if err != nil {
			return nil, err
		}
		for _, content := range v.Contents {
			fmt.Println("ObjectKey:", content.Key)
			list = append(list, FileInfo{
				Key:          content.Key,
				ETag:         content.ETag,
				Type:         1,
				Size:         content.Size,
				PartNumber:   content.PartNumber,
				LastModified: content.LastModified,
				StorageClass: content.StorageClass,
				VersionId:    content.VersionId,
			})
		}
		for _, commonPrefix := range v.CommonPrefixes {
			fmt.Println("CommonPrefixes: ", commonPrefix)
			list = append(list, FileInfo{
				Key:  commonPrefix,
				Type: 0,
			})
		}
		isTruncated = v.IsTruncated // 是否还有数据
		marker = v.NextMarker       // 设置下次请求的起始 key
	}
	return list, nil
}
