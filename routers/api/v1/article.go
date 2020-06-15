package v1

import (
	"fmt"
	"gin-demo/models"
	"gin-demo/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetArticles(c *gin.Context)  {


}

func GetArticle(c *gin.Context)  {
	idStr := c.Param("id")
	fmt.Println(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {

	}
	fmt.Println(id)
	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("无效的ID")

	code := e.INVALID_PARAMS

	var data interface{}
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			println(err.Message)
		}
	} else {
		if models.ExistArticleById(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})


}

func AddArticle(c *gin.Context)  {
	tagIdStr := c.DefaultPostForm("tag_id", "0")
	tagId, err := strconv.Atoi(tagIdStr)
	if err != nil {

	}

	stateStr := c.DefaultPostForm("state", "0")
	state, err := strconv.Atoi(stateStr)
	if err != nil {

	}
	title := c.DefaultPostForm("title", "")
	desc := c.DefaultPostForm("desc", "")
	content := c.DefaultPostForm("content", "")
	createdBy := c.DefaultPostForm("created_by", "")

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("变迁ID必须>0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("描述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0|1")

	code := e.INVALID_PARAMS

	log.Println(valid.HasErrors())

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			println(err.Message)
		}
	} else {

		log.Println("是否存在", models.ExistTagById(tagId))
		if models.ExistTagById(tagId) {

			article := models.Article{
				TagID: tagId,
				Title: title,
				Desc: desc,
				Content: content,
				CreatedBy: createdBy,
				State: state,
			}

			result := models.AddArticle(&article)
			if result {
				code = e.SUCCESS
			} else {
				code = e.ERROR_DB_INSERT
			}
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

func EditArticle(c *gin.Context)  {
	
}

func DeleteArticle(c *gin.Context)  {
	
}
