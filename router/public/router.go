package public

import (
	"Wallpaper.net/router/middleWare"
	"github.com/gin-gonic/gin"
)

/*
/api/public
*/
func Router(router *gin.RouterGroup) {
	router.GET("", middleWare.Index(" /api/public 接口首页 "))
	router.GET("/", middleWare.Index(" /api/public 接口首页 "))

	router.POST("/shell_run", RunShell)

	router.POST("/github_webhooks", GitHun)

	router.GET("/ping", middleWare.GetPing)
	router.POST("/ping", middleWare.PostPing)
}
