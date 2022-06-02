package services

func Save2cloud(uploadtype, date_format, master_uuid, logo_uuid string) string {
	lock := make(chan struct{})
	fileName := Save2tmp(uploadtype, master_uuid, logo_uuid, lock)
	go Save2minio(uploadtype, fileName, date_format, lock)
	return fileName
}
