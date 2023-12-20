/*
 * @Descripttion:tag
 * @Author: DW
 * @Date: 2023-12-18 13:39:23
 * @LastEditTime: 2023-12-19 20:42:10
 */
package controller

import (
	"blogdemo/conf"
	"blogdemo/models"

	"github.com/gin-gonic/gin"
)

func Tag(c *gin.Context) {
	tagsTemplate := models.Template.Tags

	result := models.GroupByTag(&models.ArticleList, conf.Cfg.TagDisplayQuantity)

	tagsTemplate.WriteData(c.Writer, models.BuildViewData("Tags", result))
}
