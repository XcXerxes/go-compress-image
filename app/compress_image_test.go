/*
 * @Description: 12323
 * @Author: leo
 * @Date: 2020-03-01 16:04:59
 * @LastEditors: leo
 * @LastEditTime: 2020-03-01 16:17:21
 */
package main

import (
	"testing"
)

func TestCompressImage(t *testing.T) {
	files := CompressImg("/Users/zhangjie/xinbo/myself/go/go-compress-image/")
	t.Log(files)
}
