/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-12-19 20:37:08
 * @LastEditTime: 2023-12-22 21:29:56
 */
package utils

import (
	"errors"
	"fmt"
	"strings"
)

//通过git url 返回仓库的名字
func GetRepoName(gitUrl string) (string, error) {

	if !strings.HasSuffix(gitUrl, ".git") {
		return "", errors.New("git URL must end with .git！")
	}
	fmt.Println(gitUrl)
	noSuffixUrl := strings.TrimSuffix(gitUrl, ".git")
	urlArr := strings.Split(noSuffixUrl, "/")
	fmt.Println(urlArr)
	return urlArr[len(urlArr)-1], nil
}
