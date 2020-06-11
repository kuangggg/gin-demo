package v1

import (
	"gin-demo/models"
	"gin-demo/pkg/e"
	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context)  {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	code := e.SUCCESS

	data["total"] = models.GetTagTotal(maps)
	data["lists"] = models.GetTags(0, 2, maps)

	c.JSON(200, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})

}

func AddTag(c *gin.Context)  {


}

func EditTag(c *gin.Context)  {


}

func DeleteTag(c *gin.Context)  {


}
