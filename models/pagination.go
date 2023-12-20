/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-12-18 11:34:38
 * @LastEditTime: 2023-12-18 11:41:15
 */
package models

type PageResult struct {
	List      Articles `json:"list"`
	Total     int      `json:"total"`
	Page      int      `json:"page"`
	PageSize  int      `json:"pageSize"`
	TotalPage int
}

func Pagination(articles *Articles, page int, pageSize int) PageResult {
	len := len(*articles)
	totalPage := len / pageSize

	if (len % pageSize) != 0 {
		totalPage++
	}
	result := PageResult{
		Total:     len,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
	}

	if page < 1 {
		result.Page = 1
	}

	if page > result.TotalPage {
		result.Page = result.TotalPage
	}

	if len <= result.PageSize {
		result.List = (*articles)[0:len]
	} else {
		startNum := (result.Page - 1) * result.PageSize
		endNum := startNum + result.PageSize
		if endNum > len {
			endNum = len
		}
		result.List = (*articles)[startNum:endNum]
	}
	return result
}
