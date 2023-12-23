/*
 * @Descripttion: Article
 * @Author: DW
 * @Date: 2023-12-18 13:58:17
 * @LastEditTime: 2023-12-23 17:34:36
 */
// DW
package controller

import (
	"blogdemo/models"

	"github.com/gin-gonic/gin"
)

func Article(c *gin.Context) {
	articleTemplate := models.Template.Article
	//
	if err := c.Request.ParseForm(); err != nil {
		articleTemplate.WriteError(c.Writer, err)
		return
	}
	// key
	key := c.Request.Form.Get("key")

	//  local path
	path := models.ArticleShortUrlMap[key]

	articleDetail, err := models.ReadArticleDetail(path)
	if err != nil {
		articleTemplate.WriteError(c.Writer, err)
		return
	}

	articleTemplate.WriteData(c.Writer, models.BuildViewData("Article", articleDetail))
}
