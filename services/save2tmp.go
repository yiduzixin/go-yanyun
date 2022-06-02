package services

import (
	"dropzone-go/utils"
	"os"
)

func Save2tmp(uploadtype, master_uuid, logo_uuid string, lock chan struct{}) string {
	var filename string
	if uploadtype == "image" {
		filename = image2tmp(master_uuid, logo_uuid)
	} else if uploadtype == "video" {
		filename = video2tmp(master_uuid, logo_uuid, lock)
	} else if uploadtype == "file" {
		filename = all2tmp(master_uuid)
	}
	return filename
}

func all2tmp(master_uuid string) string {
	filepath := utils.GetFile(master_uuid)
	return utils.GetName(filepath)
}

func image2tmp(master_uuid, logo_uuid string) string {
	masterFile := utils.GetFile(master_uuid)
	logoFile := utils.GetFile(logo_uuid)
	if logoFile != "" {
		utils.WaterInImage(masterFile, logoFile)
		os.Remove(logoFile)
	}
	return utils.GetName(masterFile)
}

func video2tmp(master_uuid, logo_uuid string, lock chan struct{}) string {
	masterFile := utils.GetFile(master_uuid)
	logoFile := utils.GetFile(logo_uuid)
	go utils.Video2mp4(masterFile, logoFile, lock)
	return master_uuid + ".mp4"
}
