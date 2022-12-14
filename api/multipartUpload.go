package api

import (
	"context"
	"errors"
	"fmt"
	"go-interface/client"
	"go-interface/model"

	"github.com/gin-gonic/gin"
)

type UploadMultipart struct {
	fileSystem client.FileSystem
}

type InitiateMultipartUploadRequest struct {
	Key string
}

func NewUploadMultipart(fileSystem client.FileSystem) (*UploadMultipart, error) {
	up := &UploadMultipart{
		fileSystem: fileSystem,
	}

	return up, nil
}

func (u *UploadMultipart) InitiateMultipartUpload(c *gin.Context) {
	req := &InitiateMultipartUploadRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(422, err.Error())
		return
	}

	if req.Key == "" {
		c.JSON(422, errors.New("key is null"))
		return
	}

	res, err := u.fileSystem.InitiateMultipartUpload(context.Background(), req.Key)
	if err != nil {
		c.JSON(422, err.Error())
		return
	}

	c.JSON(200, res)
}

type UploadPartRequest struct {
	Key      string `json:"key"`
	UploadId string `json:"uploadId"`
	PartNum  int32  `json:"partNum"`
	BaseMd5  string `json:"baseMd5"`
}

type UploadPartResponse struct {
	Url string `json:"url"`
}

func (u *UploadMultipart) UploadPart(c *gin.Context) {
	req := &UploadPartRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(422, err.Error())
		return
	}

	if req.Key == "" {
		c.JSON(422, errors.New("key is null"))
		return
	}

	if req.UploadId == "" {
		c.JSON(422, errors.New("upload_id is invalid"))
		return
	}

	presignedURL, err := u.fileSystem.GetPresignedURL(context.Background(), req.Key, req.UploadId, req.PartNum)
	if err != nil {
		c.JSON(422, err.Error())
		return
	}

	fmt.Println("presignedURL:", presignedURL)
	res := &UploadPartResponse{
		Url: presignedURL,
	}
	c.JSON(200, res)
}

type UploadCompleteRequest struct {
	Key      string             `json:"key"`
	UploadId string             `json:"uploadId"`
	Parts    []model.UploadList `json:"parts"`
}

func (u *UploadMultipart) UploadComplete(c *gin.Context) {
	req := &UploadCompleteRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(422, err.Error())
		return
	}

	if req.Key == "" {
		c.JSON(422, errors.New("key is null"))
		return
	}

	if req.UploadId == "" {
		c.JSON(422, errors.New("upload_id is invalid"))
		return
	}

	if err := u.fileSystem.CompleteMultipartUpload(context.Background(), req.Key, req.UploadId, req.Parts); err != nil {
		c.JSON(422, err.Error())
		return
	}

	c.JSON(200, "success")
}

type UploadListRequest struct {
	Key      string `json:"key"`
	UploadId string `json:"uploadId"`
}

func (u *UploadMultipart) UploadList(c *gin.Context) {
	req := &UploadListRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(422, err.Error())
		return
	}

	if req.Key == "" {
		c.JSON(422, errors.New("key is null"))
		return
	}

	if req.UploadId == "" {
		c.JSON(422, errors.New("upload_id is invalid"))
		return
	}

	lp, err := u.fileSystem.ListParts(context.Background(), req.UploadId, req.Key)
	if err != nil {
		c.JSON(422, err.Error())
		return
	}

	c.JSON(200, lp)
}
