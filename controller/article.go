/*
 * @Descripttion: Article
 * @Author: DW
 * @Date: 2023-12-18 13:58:17
 * @LastEditTime: 2023-12-19 10:41:34
 */
package controller

import (
	"blogdemo/models"

	"github.com/gin-gonic/gin"
)

func Article(c *gin.Context) {
	articleTemplate := models.Template.Article
	// 解析表单数据
	if err := c.Request.ParseForm(); err != nil {
		articleTemplate.WriteError(c.Writer, err)
		return
	}
	// 获取文章的键值
	key := c.Request.Form.Get("key")

	// 获取文章的路径
	path := models.ArticleShortUrlMap[key]

	articleDetail, err := models.ReadArticleDetail(path)
	if err != nil {
		articleTemplate.WriteError(c.Writer, err)
		return
	}

	articleTemplate.WriteData(c.Writer, models.BuildViewData("Article", articleDetail))
}
