package util

import (
	"gin-blog/config"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page := com.StrTo(c.Query("page")).MustInt()
	return page
}

func GetPageOffset(c *gin.Context) int {
	page := GetPage(c)
	if page > 0 {
		return (page - 1) * config.PageSize
	}
	return 0
}
