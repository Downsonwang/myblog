/*
 * @Descripttion: myblog
 * @Author: DW
 * @Date: 2023-12-17 18:00:38
 * @LastEditTime: 2023-12-17 18:05:44
 */
package route

import (
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

}
