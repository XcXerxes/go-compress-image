/*
 * @Description:123
 * @Author: leo
 * @Date: 2020-03-01 15:54:23
 * @LastEditors: leo
 * @LastEditTime: 2020-03-01 19:57:51
 */
package main

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/nfnt/resize"
)

// CompressImg 压缩图片
func CompressImg(folder string) error {
	// var (
	// 	err  error
	// 	file *os.File
	// )
	reg, _ := regexp.Compile(`^.*\.((png)|(jpg)|(jpeg))$`)
	files, _ := ioutil.ReadDir(folder)
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(f os.FileInfo) {
			name := f.Name()
			// 判断是否是图片 文件
			if reg.MatchString(name) {
				// 设置名称
				fileSuffix := path.Ext(name)
				origin := folder + name
				newName := folder + randomName() + fileSuffix
				fmt.Println(newName)

				resizeImg(origin, 0, 0, newName)
			}
			wg.Done()
		}(file)
	}
	wg.Wait()
	return nil
}

func main() {
	CompressImg("/Users/zhangjie/xinbo/myself/go/go-compress-image/")
}

// random_name 随机生成名称
func randomName() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Int())
}

// 压缩 图片
func resizePngImg(file string, width, height uint, to string) {
	// 打开图片并解码
	fileOrigin, _ := os.Open(file)
	img, _ := png.Decode(fileOrigin)
	defer fileOrigin.Close()
	m := resize.Resize(width, height, img, resize.NearestNeighbor)
	out, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	png.Encode(out, m)
}

// 压缩 图片
func resizeJpgImg(file string, width, height uint, to string) {
	// 打开图片并解码
	fileOrigin, _ := os.Open(file)
	img, _ := jpeg.Decode(fileOrigin)
	defer fileOrigin.Close()
	m := resize.Resize(width, height, img, resize.NearestNeighbor)
	out, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)
}

// Decode 图片
func resizeImg(file string, width, height uint, to string) {
	switch {
	case strings.HasSuffix(file, ".png"):
		resizePngImg(file, width, height, to)
	case strings.HasSuffix(file, ".jpg"):
		resizeJpgImg(file, width, height, to)
	default:
	}
}
