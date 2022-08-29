package qiniu

import (
	"context"
	"fmt"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var (
	accessKey = "FoTOPpjmd9O4YkvizY4UynHZbthKSdJ1TTPf9byB"
	secretKey = "-nzvMQTdT6Q5KwTZTqhu-NvsBuUBfuOQMhLQM2Vd"
	bucket    = "mhgwrite"
)

type MyPutRet struct {
	Key    string `json:"key"`
	Hash   string `json:"hash"`
	Fsize  int    `json:"fsize"`
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
	Url    string `json:"url"`
}

func Upload(key string, localFile string) (string, error) {

	key = "upload/" + key
	// 鉴权
	mac := qbox.NewMac(accessKey, secretKey)

	// 上传策略
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)","url":"http://media.mihugui.cn/$(key)"}`,
	}

	// 获取上传token
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}

	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(ret.Url)
	return ret.Url, nil

}
