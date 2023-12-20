/*
 * @Descripttion: myblog
 * @Author: DW
 * @Date: 2023-12-17 18:00:38
 * @LastEditTime: 2023-12-19 21:41:40
 */
package route

import (
	"blogdemo/conf"
	"blogdemo/controller"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", controller.Index)
	r.GET("/blog", controller.Index)
	r.GET("/categories", controller.Category)
	r.GET("/tags", controller.Tag)
	r.GET("/article", controller.Article)
	r.GET("/extra-nav", controller.ExtraNav)
	r.GET(conf.Cfg.GitHookUrl, controller.GithubHook)
	r.GET(conf.Cfg.DashboardEntrance, controller.Dashboard)

	r.Static("/public", filepath.Join(conf.Cfg.CurrentDir, "public"))
	r.Static("/assets", conf.Cfg.DocumentAssetsDir)
	r.Static("/images", filepath.Join(conf.Cfg.CurrentDir, "images"))

	return r
}
