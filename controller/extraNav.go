package controller

import (
	"blogdemo/models"

	"github.com/gin-gonic/gin"
)
// DW
func ExtraNav(c *gin.Context) {
	extraNavTemplate := models.Template.ExtraNav

	if err := c.Request.ParseForm(); err != nil {
		extraNavTemplate.WriteError(c.Writer, err)
	}

	name := c.Request.Form.Get("name")
	for _, nav := range models.Navigation {
		if nav.Title == name {
			articleDetail, err := models.ReadArticleDetail(nav.Path)
			if err != nil {
				extraNavTemplate.WriteError(c.Writer, err)
			}
			extraNavTemplate.WriteData(c.Writer, models.BuildViewData(nav.Title, articleDetail))
			return
		}
	}
}
