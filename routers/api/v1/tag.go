package v1

import (
	"fmt"
	"gin-demo/models"
	"gin-demo/pkg/e"
	"gin-demo/pkg/setting"
	"gin-demo/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)

	c.JSON(200, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})

}

func AddTag(c *gin.Context)  {

	name := c.DefaultPostForm("name", "")
	createdBy := c.DefaultPostForm("created_by", "")
	state, err := strconv.Atoi(c.DefaultPostForm("state", "0"))
	if err != nil {

	}

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长100")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态不在范围")

	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			println(err.Key)
			println(err.Message)
		}
	} else {
		if models.ExistTagByName(name) {
			code = e.ERROR_EXIST_TAG
		} else {
			code = e.SUCCESS
			models.AddTag(name, state, "kuange")
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})

}

func EditTag(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

	}

	name := c.DefaultPostForm("name", "")
	modifiedBy := c.DefaultPostForm("modified_by", "")
	state, err := strconv.Atoi(c.DefaultPostForm("state", "0"))
	if err != nil {

	}

	state, err = strconv.Atoi(c.DefaultPostForm("id", "0"))
	if err != nil {

	}


	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长100")
	valid.Required(modifiedBy, "created_by").Message("修改人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态不在范围")

	code := e.INVALID_PARAMS

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			fmt.Println(err.Key)
			fmt.Println(err.Message)
		}
	} else {

		code = e.SUCCESS

		if models.ExistTagById(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}

			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})


}

func DeleteTag(c *gin.Context)  {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

	}

	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("id 必须大于0")

	code := e.INVALID_PARAMS

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			fmt.Println(err.Message)
		}
	} else {
		code = e.SUCCESS
		if !models.ExistTagById(id) {
			code = e.ERROR_NOT_EXIST_TAG
		} else  {
			models.DeleteTag(id)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}
