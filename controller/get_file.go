package controller

import (
	"bytes"
	"dropzone-go/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

func Get_file(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	defer file.Close()
	filename := header.Filename
	_uuid := ctx.PostForm("_uuid")

	src_file := _uuid + ".uploading"
	currentdir, _ := os.Getwd()
	tmp := utils.GetTmpPath()
	tmppath := currentdir + "/" + tmp
	src_path := path.Join(tmppath, src_file)

	removing_file := tmppath + "/" + _uuid + ".removing"
	_, err = os.Stat(removing_file)
	if err == nil {
		os.Remove(removing_file)
		os.Remove(src_path)
		ctx.String(500, "Size mismatch")
	}

	current_chunk, _ := strconv.ParseInt(ctx.PostForm("dzchunkindex"), 10, 64)

	out, err := os.OpenFile(src_path, os.O_WRONLY|os.O_CREATE, 0666)

	dzchunkbyte, _ := strconv.ParseInt(ctx.PostForm("dzchunkbyteoffset"), 10, 64)
	out.Seek(dzchunkbyte, 0)
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		return
	}

	_, err = out.Write(buf.Bytes())
	if err != nil {
		return
	}

	out.Close()

	total_chunks, _ := strconv.ParseInt(ctx.PostForm("dztotalchunkcount"), 10, 64)

	if current_chunk+1 == total_chunks {
		fi, _ := os.Stat(src_path)
		dztotalfilesize, _ := strconv.ParseInt(ctx.PostForm("dztotalfilesize"), 10, 64)
		if fi.Size() != dztotalfilesize {
			ctx.String(500, "Size mismatch")
		} else {
			dst_path := path.Join(tmppath, utils.NewName(filename, _uuid))
			input, _ := ioutil.ReadFile(src_path)
			err = ioutil.WriteFile(dst_path, input, 0644)
			if err != nil {
				fmt.Println(err)
			}
			err = os.Remove(src_path)
			if err != nil {
				fmt.Println("Error creating", dst_path)
				fmt.Println(err)
				return
			}
		}
	}
}
