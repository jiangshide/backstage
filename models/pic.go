package models

import (
	"log"
	"strings"
	"os"
	"github.com/nfnt/resize"
	"image/draw"
	"time"
	"strconv"
	"io/ioutil"
	"fmt"
	"image/jpeg"
	"image/png"
	"image"
	"math/rand"
)


// 执行操作
func Cmd(path string) {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("目录" + file.Name())
			Cmd(path + file.Name() + "/")
		} else {
			if strings.Contains(strings.ToLower(file.Name()), ".jpg") {
				// 随机名称
				to := path + random_name() + ".jpg"
				origin := path + file.Name()
				fmt.Println("正在处理" + origin + ">>>" + to)
				cmd_resize(origin, 2048, 0, to)
				defer os.Remove(origin)
			}
		}
	}
}
// 改变大小
func cmd_resize(file string, width uint, height uint, to string) {
	// 打开图片并解码
	file_origin, _ := os.Open(file)
	origin, _ := jpeg.Decode(file_origin)
	defer file_origin.Close()
	canvas := resize.Resize(width, height, origin, resize.Lanczos3)
	file_out, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	defer file_out.Close()
	jpeg.Encode(file_out, canvas, &jpeg.Options{80})
	// cmd_watermark(to, strings.Replace(to, ".jpg", "@big.jpg", 1))
	cmd_thumbnail(to, 480, 360, strings.Replace(to, ".jpg", "@small.jpg", 1))
}
func cmd_thumbnail(file string, width uint, height uint, to string) {
	// 打开图片并解码
	file_origin, _ := os.Open(file)
	origin, _ := jpeg.Decode(file_origin)
	defer file_origin.Close()
	canvas := resize.Thumbnail(width, height, origin, resize.Lanczos3)
	file_out, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	defer file_out.Close()
	jpeg.Encode(file_out, canvas, &jpeg.Options{80})
}
// 水印
func cmd_watermark(file string, to string) {
	// 打开图片并解码
	file_origin, _ := os.Open(file)
	origin, _ := jpeg.Decode(file_origin)
	defer file_origin.Close()
	// 打开水印图并解码
	file_watermark, _ := os.Open("/Users/etongdai/Downloads/3606778-b978b976f8d0f591.png")
	watermark, _ := png.Decode(file_watermark)
	defer file_watermark.Close()
	//原始图界限
	origin_size := origin.Bounds()
	//创建新图层
	canvas := image.NewNRGBA(origin_size)
	// 贴原始图
	draw.Draw(canvas, origin_size, origin, image.ZP, draw.Src)
	// 贴水印图
	draw.Draw(canvas, watermark.Bounds().Add(image.Pt(30, 30)), watermark, image.ZP, draw.Over)
	//生成新图片
	create_image, _ := os.Create(to)
	jpeg.Encode(create_image, canvas, &jpeg.Options{95})
	defer create_image.Close()
}
// 随机生成文件名
func random_name() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Int())
}
