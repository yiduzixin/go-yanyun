package main

import (
	"dropzone-go/controller"
	"dropzone-go/libs"
	"dropzone-go/utils"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	currentdir, _ := os.Getwd()
	tmp := utils.GetTmpPath()
	tmppath := currentdir + "/" + tmp
	_, err := os.Stat(tmppath)
	if err != nil {
		err := os.MkdirAll(tmppath, os.ModePerm)
		if err != nil {
			return
		}
	}

	r := gin.Default()
	r.NoMethod(libs.HandleNotFound)
	r.NoRoute(libs.HandleNotFound)
	r.LoadHTMLGlob("templates/*")
	r.Static("static", "static")
	submit := r.Group("/")

	{
		submit.GET("/", controller.ShowIndex)
		submit.POST("/", libs.CatchE(controller.GetFormData))
		submit.POST("/getfile", controller.Get_file)
	}

	err = r.Run(":8000")
	if err != nil {
		return
	}
}
