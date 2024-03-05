/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-12-17 17:34:16
 * @LastEditTime: 2024-03-04 19:15:24
 */
package controller

import (
	"blogdemo/config"
	"blogdemo/models"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {

	indexTemplate := models.Template.Index

	if err := r.ParseForm(); err != nil {
		indexTemplate.WriteError(w, err)
	}

	page, err := strconv.Atoi(r.Form.Get("page"))
	if err != nil {
		page = 1
	}
	articles := models.ArticleList

	search := r.Form.Get("search")
	category := r.Form.Get("category")
	tag := r.Form.Get("tag")

	if search != "" || category != "" || tag != "" {
		articles = models.ArticleSearch(&articles, search, category, tag)
	}

	result := models.Pagination(&articles, page, config.Cfg.PageSize)

	indexTemplate.WriteData(w, models.BuildViewData("Blog", result))
}
