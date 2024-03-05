/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-12-17 17:34:16
 * @LastEditTime: 2024-03-04 19:11:19
 */
package main

import (
	"blogdemo/config"
	"blogdemo/models"
	"blogdemo/routes"
	"fmt"
	"net/http"
	"strconv"
)

func init() {
	models.CompiledContent() //克隆或者更新文章、递归生成文章、导航、短链 Map、加载模板
}

func main() {

	routes.InitRoute()
	fmt.Printf("Version：v%v \n", config.Cfg.Version)
	fmt.Printf("ListenAndServe On Port %v \n", config.Cfg.Port)
	fmt.Printf("UpdateArticle's GitHookUrl: %v   Secret:  %v \n", config.Cfg.GitHookUrl, config.Cfg.WebHookSecret)
	if err := http.ListenAndServe(":"+strconv.Itoa(config.Cfg.Port), nil); err != nil {
		fmt.Println("ServeErr:", err)
	}
}
