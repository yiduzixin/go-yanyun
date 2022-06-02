package utils

import (
	"github.com/Unknwon/goconfig"
	"os"
)

func GetAccessInfo() (ACCESS_KEY string, SECRET_KEY string) {
	currentdir, _ := os.Getwd()
	filepath := currentdir + "/config.ini"
	cfg, _ := goconfig.LoadConfigFile(filepath)
	ACCESS_KEY, _ = cfg.GetValue("accessinfo", "ACCESS_KEY")
	SECRET_KEY, _ = cfg.GetValue("accessinfo", "SECRET_KEY")
	return ACCESS_KEY, SECRET_KEY
}

func GetTmpPath() (Tmp string) {
	currentdir, _ := os.Getwd()
	filepath := currentdir + "/config.ini"
	cfg, _ := goconfig.LoadConfigFile(filepath)
	Tmp, _ = cfg.GetValue("os_env", "TMP_FOLDER")
	return Tmp
}

func GetProxyHost() string {
	currentdir, _ := os.Getwd()
	filepath := currentdir + "/config.ini"
	cfg, _ := goconfig.LoadConfigFile(filepath)
	host, _ := cfg.GetValue("imageproxy", "HOST")
	return host
}

func GetMinioBacket() string {
	currentdir, _ := os.Getwd()
	filepath := currentdir + "/config.ini"
	cfg, _ := goconfig.LoadConfigFile(filepath)
	Backet, _ := cfg.GetValue("minio", "BUCKET")
	return Backet
}

func GetMinioHost() string {
	currentdir, _ := os.Getwd()
	filepath := currentdir + "/config.ini"
	cfg, _ := goconfig.LoadConfigFile(filepath)
	host, _ := cfg.GetValue("minio", "MINIO_HOST")
	return host
}

func GetMinioEndpoint() string {
	currentdir, _ := os.Getwd()
	filepath := currentdir + "/config.ini"
	cfg, _ := goconfig.LoadConfigFile(filepath)
	Endpoint, _ := cfg.GetValue("minio", "MINIO_ENDPOINT")
	return Endpoint
}

func GetMinioAccessKey() string {
	currentdir, _ := os.Getwd()
	filepath := currentdir + "/config.ini"
	cfg, _ := goconfig.LoadConfigFile(filepath)
	AccessKey, _ := cfg.GetValue("minio", "MINIO_ACCESS_KEY")
	return AccessKey
}

func GetMinioSecretKey() string {
	currentdir, _ := os.Getwd()
	filepath := currentdir + "/config.ini"
	cfg, _ := goconfig.LoadConfigFile(filepath)
	SecretKey, _ := cfg.GetValue("minio", "MINIO_SECRET_KEY")
	return SecretKey
}

func GetMiniouseSSL() bool {
	currentdir, _ := os.Getwd()
	filepath := currentdir + "/config.ini"
	cfg, _ := goconfig.LoadConfigFile(filepath)
	useSSL, _ := cfg.GetValue("minio", "MINIO_useSSL")
	if useSSL == "true" {
		return true
	}
	return false
}
