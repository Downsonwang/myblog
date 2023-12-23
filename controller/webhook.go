/*
 * @Descripttion: githubhook
 * @Author:DW
 * @Date: 2023-12-18 20:51:06
 * @LastEditTime: 2023-12-22 22:46:12
 */
package controller

// DW
import (
	"blogdemo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GithubHook(c *gin.Context) {
	SedResponse(c, http.StatusOK, "ok")
	models.CompiledContent()
}

func SedResponse(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"msg": msg})
}
