package utils

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func WaterInImage(src_img, logo string) {
	img_file, _ := os.Open(src_img)
	extName := ExtName(src_img)
	var img image.Image
	if extName == ".jpg" || extName == ".jpeg" {
		img, _ = jpeg.Decode(img_file)
	} else if extName == ".png" {
		img, _ = png.Decode(img_file)
	} else {
		fmt.Println("暂不支持文件类型")
		return
	}

	wmb_file, _ := os.Open(logo)
	defer wmb_file.Close()
	wmb_img, _ := png.Decode(wmb_file)

	offset := image.Pt(img.Bounds().Dx()-wmb_img.Bounds().Dx()-10, img.Bounds().Dy()-wmb_img.Bounds().Dy()-10)
	b := img.Bounds()
	m := image.NewRGBA(b)

	draw.Draw(m, b, img, image.Point{0, 0}, draw.Src)
	draw.Draw(m, wmb_img.Bounds().Add(offset), wmb_img, image.Point{0, 0}, draw.Over)

	img_file.Close()
	imgw, _ := os.Create(src_img)

	if extName == ".jpg" || extName == ".jpeg" {
		jpeg.Encode(imgw, m, &jpeg.Options{100})
	} else {
		png.Encode(imgw, m)
	}
	imgw.Close()
}
