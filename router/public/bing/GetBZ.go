package bing

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"Wallpaper.net/global/config"
	"github.com/gin-gonic/gin"
)

func GetBZ(c *gin.Context) {
	idx := c.DefaultQuery("idx", "0")
	fileUrl := getFileName(idx, config.BingImg)

	fmt.Println(c.FullPath())

	c.Redirect(http.StatusMovedPermanently, fileUrl)
}

func getFileName(str string, fileArr []config.BingImgType) string {
	num, _ := strconv.Atoi(str) // 错误的值会返回 0
	max := len(fileArr)
	rNum := num - 1
	if num < 1 {
		rand.Seed(time.Now().Unix())
		rNum = rand.Intn(len(fileArr))
	}
	if num > max {
		rNum = max
	}
	return fileArr[rNum].Path
}
