/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-12-18 12:15:41
 * @LastEditTime: 2023-12-22 22:45:44
 */
// DW
package controller

import (
	"blogdemo/conf"
	"blogdemo/models"

	"github.com/gin-gonic/gin"
)

func Category(c *gin.Context) {
	categoriesTemplate := models.Template.Categories
	result := models.GroupByCategory(&models.ArticleList, conf.Cfg.CategoryDisplayQuantity)

	categoriesTemplate.WriteData(c.Writer, models.BuildViewData("Categories", result))

}
