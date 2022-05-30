package public

import (
	"Wallpaper.net/router/middleWare"
	"Wallpaper.net/router/public/bing"
	"github.com/gin-gonic/gin"
)

/*
/api/public
*/
func Router(router *gin.RouterGroup) {
	router.GET("", middleWare.Index(" /api/public 接口首页 "))
	router.GET("/", middleWare.Index(" /api/public 接口首页 "))

	router.GET("/url", bing.GetUrl)
	router.GET("/bz", bing.GetBZ)

	router.GET("/ping", middleWare.GetPing)
	router.POST("/ping", middleWare.PostPing)
}
