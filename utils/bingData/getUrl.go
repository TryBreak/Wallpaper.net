package bingData

import (
	"fmt"
	"strings"
	"time"

	"Wallpaper.net/global"
	"Wallpaper.net/global/config"
	"github.com/EasyGolang/goTools/mFetch"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

const BaseUrl = "//bing.com"

type ImagesType struct {
	Startdate     string `json:"startdate"`
	Enddate       string `json:"enddate"`
	URL           string `json:"url"`
	Copyright     string `json:"copyright"`
	Copyrightlink string `json:"copyrightlink"`
	Title         string `json:"title"`
}

type ResType struct {
	Images []ImagesType `json:"images"`
}

func Start() {
	ImgDataPath := "/HPImageArchive.aspx?format=js&n=8"

	reqData := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: "http:" + BaseUrl,
		Path:   ImgDataPath,
	}).Get()

	if len(reqData) < 5 {
		errStr := fmt.Errorf("接口请求失败: %+v", reqData)
		global.LogErr(errStr)
		return
	}

	var data ResType
	err := jsoniter.Unmarshal(reqData, &data)
	if err != nil {
		errStr := fmt.Errorf("数据格式化失败 : " + mStr.ToStr(reqData))
		global.LogErr(mStr.ToStr(errStr))
		return
	}

	imageArr := data.Images

	var images []config.BingImgType

	for _, v := range imageArr {
		startTime, err := time.ParseInLocation("20060102", v.Startdate, time.Local)
		if err != nil {
			continue
		}
		endTime, err := time.ParseInLocation("20060102", v.Enddate, time.Local)
		if err != nil {
			continue
		}

		FullUrl := BaseUrl + v.URL
		UrlArr := strings.Split(FullUrl, "&")
		if len(UrlArr) > 0 {
			FullUrl = UrlArr[0]
		}

		img := config.BingImgType{
			StartTime:     mTime.ToUnixMsec(startTime),
			EndTime:       mTime.ToUnixMsec(endTime),
			Path:          FullUrl,
			Url:           "https:" + FullUrl,
			Copyright:     v.Copyright,
			CopyrightLink: v.Copyrightlink,
			Title:         v.Title,
		}

		images = append(images, img)
	}

	config.BingImg = images
}
