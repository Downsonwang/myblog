/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-12-18 10:33:25
 * @LastEditTime: 2023-12-22 22:46:06
 */
package controller
// DW
import (
	"blogdemo/conf"
	"blogdemo/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	indexTemp := models.Template.Index
	if err := c.Request.ParseForm(); err != nil {
		indexTemp.WriteError(c.Writer, err)
		return
	}

	page, err := strconv.Atoi(c.Request.Form.Get("page"))
	if err != nil {
		page = -1
	}
	articles := models.ArticleList

	search := c.Request.Form.Get("search")
	category := c.Request.Form.Get("category")
	tag := c.Request.Form.Get("tag")

	// 根据搜索条件过滤文章列表
	if search != "" || category != "" || tag != "" {
		articles = models.ArticleSearch(&articles, search, category, tag)
	}

	result := models.Pagination(&articles, page, conf.Cfg.PageSize)
	indexTemp.WriteData(c.Writer, models.BuildViewData("Blog", result))
}
