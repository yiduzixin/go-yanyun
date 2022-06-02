package services

import (
	"dropzone-go/utils"
	"fmt"
	"github.com/minio/minio-go/v6"
	"os"
	"path"
)

func getFolder(uploadtype string) string {
	typeFolder := make(map[string]string)
	typeFolder["image"] = "images"
	typeFolder["video"] = "video"
	typeFolder["file"] = "files"

	return typeFolder[uploadtype]
}

func Save2minio(uploadtype, file_name, date_format string, lock chan struct{}) {

	folder := getFolder(uploadtype)
	currentdir, _ := os.Getwd()
	srcpath := path.Join(currentdir, utils.GetTmpPath(), file_name)
	dstpath := folder + "/" + date_format + "/" + file_name

	defer os.Remove(srcpath)

	//获取minio信息
	backet := utils.GetMinioBacket()
	endpoint := utils.GetMinioEndpoint()
	AccessKey := utils.GetMinioAccessKey()
	SecretKey := utils.GetMinioSecretKey()
	useSSL := utils.GetMiniouseSSL()

	minioClient, err := minio.New(endpoint, AccessKey, SecretKey, useSSL)
	if err != nil {
		fmt.Println(err)
	}

	//等待视频转换完成，再上传到minio
	if uploadtype == "video" {
		select {
		case <-lock:
			minioClient.FPutObject(backet, dstpath, srcpath, minio.PutObjectOptions{ContentType: "application/octet-stream"})
			return
		}
	}

	_, err = minioClient.FPutObject(backet, dstpath, srcpath, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
	}

}
