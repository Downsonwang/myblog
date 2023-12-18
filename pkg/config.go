/*
 * @Descripttion:  read config
 * @Author: DW
 * @Date: 2023-12-17 20:46:08
 * @LastEditTime: 2023-12-17 21:33:52
 */
package pkg

import (
	"blogdemo/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	models.HomeInfo
	models.SystemInfo
}

var Cfg Config

func init() {

	if dir, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		Cfg.CurrentDir = dir
	}

	pathFile, err := ioutil.ReadFile(Cfg.CurrentDir + "/config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(pathFile, &Cfg)
	if err != nil {
		panic(err)
	}

	if Cfg.DashboardEntrance == "" || !strings.HasPrefix(Cfg.DashboardEntrance, "/") {
		Cfg.DashboardEntrance = "/admin"
	}

	repoName, err := GetRepoName(Cfg.DocumentGitUrl)
	if err != nil {
		panic(err)
	}

	Cfg.AppName = "Boke"
	Cfg.Version = 1.0
	Cfg.DocumentDir = Cfg.CurrentDir + "/" + repoName
	Cfg.GitHookUrl = "/api/get_hook"
	Cfg.AppRepository = "https://github.com/Downsonwang/myblog.git"
}

func GetRepoName(gitUrl string) (string, error) {
	if !strings.HasSuffix(gitUrl, ".git") {
		return "", errors.New("git Url must end with .git")
	}
	noSuffixUrl := strings.TrimSuffix(gitUrl, ".git")
	urlArr := strings.Split(noSuffixUrl, "/")
	return urlArr[len(urlArr)-1], nil
}

func CheckInit() {
	if _, err := exec.LookPath("git"); err != nil {
		fmt.Println("请先安装git")
		panic(err)
	}
	if !IsDir(Cfg.DocumentDir) {
		fmt.Println("正在克隆文档仓库，请稍等...")
		out, err := RunCmdByDir(Cfg.CurrentDir, "git", "clone", Cfg.DocumentGitUrl)
		if err != nil {
			panic(err)
		}
		fmt.Println(out)
	} else {
		out, err := RunCmdByDir(Cfg.DocumentDir, "git", "pull")
		fmt.Println(out)
		if err != nil {
			panic(err)
		}

	}
	if err := checkDocDirAndBindConfig(&Cfg); err != nil {
		fmt.Println("文档缺少必要的目录")
		panic(err)
	}

	imgDir := Cfg.CurrentDir + "/images"
	if !IsDir(imgDir) {
		if os.Mkdir(imgDir, os.ModePerm) != nil {
			panic("生成images目录失败！")
		}
	}
}
