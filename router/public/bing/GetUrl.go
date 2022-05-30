package bing

import (
	"net/http"

	"Wallpaper.net/global/config"
	"Wallpaper.net/router/ginResult"
	"github.com/gin-gonic/gin"
)

func GetUrl(c *gin.Context) {
	c.JSON(http.StatusOK, ginResult.OK.WithData(config.BingImg))
}
