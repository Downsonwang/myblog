/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-12-17 22:21:32
 * @LastEditTime: 2023-12-19 22:17:56
 */
package models

import (
	"blogdemo/conf"
	"blogdemo/utils"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
)

type Article struct {
	Title    string   `json:"title"`
	Desc     string   `json:"Desc"`
	Tags     []string `json:"tags"`
	Author   string   `json:"author"`
	Path     string
	ShortUrl string
	Category string
	Date     time.Time `json:"date"`
}

type Articles []Article
type ArticleDetail struct {
	Article
	Body string
}

var ArticleList Articles

type Category struct {
	Name     string
	Quantity int
	Articles Articles
}
type Categories []Category

type Tag struct {
	Name     string
	Quantity int
	Articles Articles
}

type Tags []Tag

func ArticleSearch(articles *Articles, search, category, tag string) Articles {
	var articleList Articles
	for _, article := range *articles {
		pass := true
		if search != "" && !strings.Contains(article.Title, search) {
			pass = false
		}
		if category != "" && !strings.Contains(article.Category, category) {
			pass = false
		}

		if tag != "" {
			pass = false
			for _, tagItem := range article.Tags {
				if !strings.Contains(tagItem, tag) {
					pass = true
					break
				}
			}
		}
		if pass {
			articleList = append(articleList, article)
		}
	}
	return articleList
}

func RecursiveReadArticles(dir string) (Articles, error) {

	var articles Articles

	dirInfo, err := os.Stat(dir)

	if err != nil {
		return articles, err
	}
	if !dirInfo.IsDir() {
		return articles, errors.New("目标不是一个目录")
	}

	fileOrDir, err := ioutil.ReadDir(dir)
	if err != nil {
		return articles, err
	}

	for _, fileInfo := range fileOrDir {
		name := fileInfo.Name()
		path := dir + "/" + name
		upperName := strings.ToUpper(name)
		if fileInfo.IsDir() {
			subArticles, err := RecursiveReadArticles(path)
			if err != nil {
				return articles, err
			}
			articles = append(articles, subArticles...)
		} else if strings.HasSuffix(upperName, ".MD") {
			article, err := ReadArticle(path)
			if err != nil {
				return articles, err
			}
			articles = append(articles, article)
		} else if strings.HasSuffix(upperName, ".PNG") ||
			strings.HasSuffix(upperName, ".GIF") ||
			strings.HasSuffix(upperName, ".JPG") {

			dst := conf.Cfg.CurrentDir + "/images/" + name
			if !utils.IsFile(dst) {
				_, _ = utils.CopyFile(path, dst)
			}
		}

	}
	return articles, nil
}

func ReadArticle(path string) (Article, error) {
	article, _, err := readMarkdown(path)
	if err != nil {
		return article, err
	}
	return article, nil
}

func GetCategoryName(path string) string {
	var categoryName string
	newPath := strings.Replace(path, conf.Cfg.DocumentContentDir+"/", "", 1)

	if !strings.Contains(newPath, "/") { //文件在根目录下(content/)没有分类名称
		categoryName = "未分类"
	} else {
		categoryName = strings.Split(newPath, "/")[0]
	}
	return categoryName
}

func GroupByCategory(articles *Articles, articleQuantity int) Categories {
	var categories Categories
	categoryMap := make(map[string]Articles)

	for _, article := range *articles {
		_, existedCategory := categoryMap[article.Category]
		if existedCategory {
			categoryMap[article.Category] = append(categoryMap[article.Category], article)
		} else {
			categoryMap[article.Category] = Articles{article}
		}
	}
	for categoryName, articles := range categoryMap {
		articleLen := len(articles)

		var articleList Articles
		if articleQuantity <= 0 {
			articleList = articles
		} else {
			if articleQuantity > articleLen {
				articleList = articles[0:articleLen]
			} else {
				articleList = articles[0:articleQuantity]
			}
		}
		categories = append(categories, Category{
			Name:     categoryName,
			Quantity: articleLen,
			Articles: articleList,
		})
	}
	sort.Sort(categories)
	return categories
}

func (c Categories) Len() int { return len(c) }

func (c Categories) Less(i, j int) bool { return c[i].Quantity > c[j].Quantity }

func (c Categories) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

func GroupByTag(articles *Articles, articleQuantity int) Tags {
	var tags Tags
	tagMap := make(map[string]Articles)
	for _, article := range *articles {
		for _, tag := range article.Tags {
			_, exsitedCategory := tagMap[tag]
			if exsitedCategory {
				tagMap[tag] = append(tagMap[tag], article)
			} else {
				tagMap[tag] = Articles{article}
			}
		}
	}

	for categoryName, articleItem := range tagMap {
		articleLen := len(articleItem)
		var articleList Articles
		if articleQuantity <= 0 {
			articleList = articleItem
		} else {
			if articleQuantity > articleLen {
				articleList = articleItem[0:articleLen]
			} else {
				articleList = articleItem[0:articleQuantity]
			}
		}
		tags = append(tags, Tag{
			Name:     categoryName,
			Quantity: articleLen,
			Articles: articleList,
		})
	}
	sort.Sort(tags)
	return tags
}

func (c Tags) Len() int { return len(c) }

func (c Tags) Less(i, j int) bool { return c[i].Quantity > c[j].Quantity }

func (c Tags) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

func ReadArticleDetail(path string) (ArticleDetail, error) {
	_, articleDetail, err := readMarkdown(path)
	if err != nil {
		return articleDetail, err
	}
	return articleDetail, nil
}

func readMarkdown(path string) (Article, ArticleDetail, error) {
	var article Article
	var articleDetail ArticleDetail
	mdFile, err := os.Stat(path)
	if err != nil {
		return article, articleDetail, err
	}
	if mdFile.IsDir() {
		return article, articleDetail, errors.New("this path is Dir")
	}
	markdown, err := ioutil.ReadFile(path)
	if err != nil {
		return article, articleDetail, err
	}
	markdown = bytes.TrimSpace(markdown)

	article.Path = path
	article.Category = GetCategoryName(path)
	article.Title = strings.TrimSuffix(strings.ToUpper(mdFile.Name()), ".MD")
	article.Date = time.Time(mdFile.ModTime())

	if !bytes.HasPrefix(markdown, []byte("```json")) {
		article.Desc = cropDesc(markdown)
		articleDetail.Article = article
		articleDetail.Body = string(markdown)
		return article, articleDetail, nil
	}
	markdown = bytes.Replace(markdown, []byte("```json"), []byte(""), 1)
	markdownArrInfo := bytes.SplitN(markdown, []byte("```"), 2)
	article.Desc = cropDesc(markdownArrInfo[1])

	if err := json.Unmarshal(bytes.TrimSpace(markdownArrInfo[0]), &article); err != nil {
		article.Title = "文章[" + article.Title + "]解析 JSON 出错，请检查。"
		article.Desc = err.Error()
		return article, articleDetail, nil
	}
	article.Path = path
	article.Title = strings.ToUpper(article.Title)

	articleDetail.Article = article

	var buf bytes.Buffer
	if err := goldmark.Convert(markdownArrInfo[1], &buf); err != nil {
		article.Title = "文章[" + article.Title + "]解析 markdown 出错，请检查。"
		return article, articleDetail, nil
	}

	articleDetail.Body = buf.String()
	return article, articleDetail, nil
}

func cropDesc(c []byte) string {
	content := []rune(string(c))
	contentLen := len(content)

	if contentLen <= conf.Cfg.DescriptionLen {
		return string(content[0:contentLen])
	}

	return string(content[0:conf.Cfg.DescriptionLen])
}

func (a Articles) Len() int { return len(a) }

func (a Articles) Less(i, j int) bool { return time.Time(a[i].Date).After(time.Time(a[j].Date)) }

func (a Articles) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func initArticlesAndImages(dir string) (Articles, map[string]string, error) {
	var articles Articles
	shortUrlMap := make(map[string]string)

	articles, err := RecursiveReadArticles(dir)
	if err != nil {
		return articles, shortUrlMap, err
	}
	sort.Sort(articles)
	for i := len(articles) - 1; i >= 0; i-- {
		//这里必须使用倒序的方式生成 shortUrl,因为如果有相同的文章标题，
		// 倒序会将最老的文章优先生成shortUrl，保证和之前的 shortUrl一样
		article := articles[i]
		keyword := utils.GenerateShortUrl(article.Title, func(url, keyword string) bool {
			//保证 keyword 唯一
			_, ok := shortUrlMap[keyword]
			return !ok
		})
		articles[i].ShortUrl = keyword
		shortUrlMap[keyword] = article.Path
	}
	return articles, shortUrlMap, nil
}
