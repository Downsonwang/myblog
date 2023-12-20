/*
 * @Descripttion: template
 * @Author: DW
 * @Date: 2023-12-18 10:37:51
 * @LastEditTime: 2023-12-19 20:35:19
 */
package models

/* config cfg*/

import (
	"blogdemo/conf"
	"fmt"
	"html/template"
	"io"
)

type TemplatePointer struct {
	*template.Template
}

type HtmlTemplate struct {
	Article    TemplatePointer
	Categories TemplatePointer
	Tags       TemplatePointer
	Dashboard  TemplatePointer
	ExtraNav   TemplatePointer
	Index      TemplatePointer
}

var Template HtmlTemplate

func (t TemplatePointer) WriteData(w io.Writer, data interface{}) {
	// 渲染模版将数据填充
	err := t.Execute(w, data)
	if err != nil {
		if _, e := w.Write([]byte(err.Error())); e != nil {
			fmt.Println(e)
		}
	}
}

func (t TemplatePointer) WriteError(w io.Writer, err error) {
	if _, e := w.Write([]byte(err.Error())); e != nil {
		fmt.Println(e)
	}

}

func initHtmlTemplate(viewDir string) (HtmlTemplate, error) {
	var htmlTemplate HtmlTemplate
	var err error

	if htmlTemplate.Index, err = readHtmlTemplate("index", viewDir); err != nil {
		return htmlTemplate, err
	}
	if htmlTemplate.ExtraNav, err = readHtmlTemplate("extraNav", viewDir); err != nil {
		return htmlTemplate, err
	}
	if htmlTemplate.Dashboard, err = readHtmlTemplate("dashboard", viewDir); err != nil {
		return htmlTemplate, err
	}
	if htmlTemplate.Categories, err = readHtmlTemplate("categories", viewDir); err != nil {
		return htmlTemplate, err
	}
	if htmlTemplate.Article, err = readHtmlTemplate("article", viewDir); err != nil {
		return htmlTemplate, err
	}
	if htmlTemplate.Tags, err = readHtmlTemplate("tags", viewDir); err != nil {
		return htmlTemplate, err
	}

	return htmlTemplate, nil
}

func SpreadDigit(n int) []int {
	var r []int
	for i := 1; i <= n; i++ {
		r = append(r, i)
	}
	return r
}
func readHtmlTemplate(htmlFileName string, viewDir string) (TemplatePointer, error) {

	head := viewDir + "/layouts/head.html"
	footer := viewDir + "/layouts/footer.html"

	tp, err := template.New(htmlFileName+".html").
		Funcs(template.FuncMap{"SpreadDigit": SpreadDigit}).
		ParseFiles(viewDir+"/"+htmlFileName+".html", head, footer)
	if err != nil {
		return TemplatePointer{}, err
	}
	return TemplatePointer{tp}, nil
}

func BuildViewData(title string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"Title":  title,
		"Data":   data,
		"Config": conf.Cfg,
		"Navs":   Navigation,
	}
}
