package client

import (
	"context"
	"fmt"
	"go-interface/config"
	"go-interface/model"
	"os"
	"path"
	"path/filepath"
	"testing"
)

// 一个完整的上传测试用例
func TestFileSystem(t *testing.T) {
	conf := config.DefaultTestConfig()
	fileSystem, err := NewFileSystem(conf)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	name := "d/ab.jpg"
	filePath := "../file/axx.jpg"

	if !path.IsAbs(filePath) {
		filePath, err = filepath.Abs(filePath)
		if err != nil {
			t.Fatal(err)
		}
	}

	initInfo, err := fileSystem.InitiateMultipartUpload(ctx, name)
	if err != nil {
		t.Fatal(err)
	}

	url, err := fileSystem.GetPresignedURL(ctx, name, initInfo.UploadId, 1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("url", url)
	// 通过预签名方式上传对象
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}
	// 本地文件可以通过两种上传方式，http 请求，或者SDK
	// 1、SDK
	etag, err := fileSystem.UploadPart(ctx, initInfo.UploadId, name, "", 1, data)
	if err != nil {
		t.Fatal(err)
	}

	// 2、Http 请求
	// f := bytes.NewReader(data)
	// req, err := http.NewRequest(http.MethodPut, url, f)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// // 用户可自行设置请求头部
	// req.Header.Set("Content-Type", "image/jpeg")
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Printf("resp:%#v\n", resp)
	// etag := resp.Header.Get("Etag")

	parts := make([]model.UploadList, 0)
	parts = append(parts, model.UploadList{
		PartNum: 1,
		// Etag:    resp.Header.Get("Etag"),
		Etag: etag,
	})

	if err := fileSystem.CompleteMultipartUpload(ctx, initInfo.UploadId, name, parts); err != nil {
		t.Fatal(err)
	}
}
