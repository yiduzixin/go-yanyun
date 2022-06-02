package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

func Video2mp4(videoPath, logo string, lock chan struct{}) {
	dir := path.Dir(videoPath)
	baseName := path.Base(videoPath)
	extname := path.Ext(baseName)
	suffixName := strings.TrimSuffix(baseName, extname)
	inputPath := videoPath
	outputPath := path.Join(dir, suffixName+".mp4")
	mp4s := []string{".mp4", ".MP4", ".Mp4", ".mP4"}
	for _, mp4 := range mp4s {
		if mp4 == extname {
			dst_path := path.Join(dir, suffixName+".flv")
			inputPath = dst_path
			input, err := ioutil.ReadFile(videoPath)
			if err != nil {
				fmt.Println(err)
			}
			err = ioutil.WriteFile(dst_path, input, 0644)
			if err != nil {
				fmt.Println(err)
			}
			_ = os.Remove(videoPath)
		}
	}
	waterParam := "movie=./" + GetTmpPath() + "/" + GetName(logo) + " [watermark]; [in][watermark] overlay=main_w-overlay_w-10:main_h-overlay_h-10 [out]"
	cmdArguments := []string{"-i", inputPath, outputPath}
	if logo != "" {
		cmdArguments = []string{"-i", inputPath, "-vf", waterParam, outputPath}
	}
	cmd := exec.Command("ffmpeg", cmdArguments...)
	cmd.Run()

	if logo != "" {
		os.Remove(logo)
	}
	os.Remove(inputPath)
	lock <- struct{}{}
	close(lock)
}
