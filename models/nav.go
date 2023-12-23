/*
 * @Descripttion: 更新文档库生成内容
 * @Author:DW
 * @Date: 2023-12-18 10:57:19
 * @LastEditTime: 2023-12-19 20:46:04
 */
package models

import (
	"blogdemo/conf"
	"sort"
	"strings"
	"sync"
)

type Nav struct {
	Title string
	Path  string
}

type Navs []Nav

var Navigation Navs
var ArticleShortUrlMap map[string]string //用来保证文章 shortUrl 唯一和快速定位文章

func initExtraNav(dir string) (Navs, error) {
	var navigation Navs
	var extraNav Articles

	extraNav, err := RecursiveReadArticles(dir)
	if err != nil {
		return navigation, err
	}
	sort.Sort(extraNav)

	for _, article := range extraNav {
		title := strings.Title(strings.ToLower(article.Title))
		navigation = append(navigation, Nav{Title: title, Path: article.Path})
	}

	return navigation, nil
}

func CompiledContent() {
	conf.CheckInit() //克隆或者更新文档库
	//下面是对内容的生成
	wg := sync.WaitGroup{}
	var err error
	//导航
	wg.Add(1)
	go func() {
		Navigation, err = initExtraNav(conf.Cfg.DocumentExtraNavDir)
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()

	//加载html模板
	wg.Add(1)
	go func() {
		Template, err = initHtmlTemplate(conf.Cfg.CurrentDir + "/views")
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()

	//文章
	wg.Add(1)
	go func() {
		ArticleList, ArticleShortUrlMap, err = initArticlesAndImages(conf.Cfg.DocumentContentDir)
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()
	wg.Wait()
	//启用并发比之前节约4倍左右的时间
}
