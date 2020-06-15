package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(c *gin.Context) int  {
	pageStr := c.DefaultQuery("page", "0")

	if page, err := strconv.Atoi(pageStr); err == nil {
		return page
	}

	return 0
}




