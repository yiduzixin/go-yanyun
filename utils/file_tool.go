package utils

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func GetFile(name_uuid string) string {
	currentdir, _ := os.Getwd()
	dir := path.Join(currentdir, GetTmpPath())
	files, _ := ioutil.ReadDir(dir)
	filename := ""
	for _, file := range files {
		baseName := path.Base(file.Name())
		extname := path.Ext(baseName)
		if strings.TrimSuffix(baseName, extname) == name_uuid {
			filename = baseName
		}
	}
	if filename == "" {
		return ""
	}
	filepath := path.Join(dir, filename)
	return filepath
}

func GetName(filepath string) string {
	_, filename := path.Split(filepath)
	return filename
}

func NewName(filename, newname string) string {
	extName := ExtName(filename)
	if extName != "" {
		return newname + extName
	}
	return newname
}

func ExtName(filename string) string {
	return path.Ext(filename)
}
