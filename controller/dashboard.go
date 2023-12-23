/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-12-19 10:51:43
 * @LastEditTime: 2023-12-22 22:45:49
 */
package controller

import (
	"blogdemo/conf"
	"blogdemo/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DW
func Dashboard(c *gin.Context) {

	var dashboardMsg []string
	dashboardTemplate := models.Template.Dashboard

	if err := c.Request.ParseForm(); err != nil {
		dashboardTemplate.WriteError(c.Writer, err)
	}

	index, err := strconv.Atoi(c.Request.Form.Get("theme"))
	if err == nil && index < len(conf.Cfg.ThemeOption) {
		conf.Cfg.ThemeColor = conf.Cfg.ThemeOption[index]
		dashboardMsg = append(dashboardMsg, "颜色切换成功!")
	}

	action := c.Request.Form.Get("action")
	if "updateArticle" == action {
		models.CompiledContent()
		dashboardMsg = append(dashboardMsg, "文章更新成功!")
	}

	dashboardTemplate.WriteData(c.Writer, models.BuildViewData("Dashboard", map[string]interface{}{
		"msg": dashboardMsg,
	}))

}
