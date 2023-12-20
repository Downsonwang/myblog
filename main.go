/*
 * @Descripttion:  BLOG
 * @Author:  DW
 * @Date: 2023-12-17 17:45:17
 * @LastEditTime: 2023-12-20 15:40:18
 */
package main

import (
	"blogdemo/conf"
	"blogdemo/models"
	"blogdemo/route"
	"fmt"
	"net/http"
)

func init() {
	models.CompiledContent() ////克隆或者更新文章、递归生成文章、导航、短链 Map、加载模板
}

func main() {
	r := route.InitRoute()

	s := &http.Server{
		Addr:    "154.8.202.7:" + fmt.Sprintf("%d", conf.Cfg.Port),
		Handler: r,
	}
	if err := s.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
