package model

type UploadList struct {
	PartNum int    `json:"partNum"`
	Etag    string `json:"etag"`
}

type InitiateMultipartUploadResult struct {
	Bucket    string
	Key       string
	UploadId  string `json:"uploadId"`
	RequestId string `json:"requestId,omitempty"`
}
