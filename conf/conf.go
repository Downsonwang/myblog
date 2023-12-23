package conf

import (
	"blogdemo/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	HomeInfo
	SystemInfo
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
	//fmt.Println(pathFile)
	err = json.Unmarshal(pathFile, &Cfg)
	if err != nil {
		panic(err)
	}

	if Cfg.DashboardEntrance == "" || !strings.HasPrefix(Cfg.DashboardEntrance, "/") {
		Cfg.DashboardEntrance = "/admin"
	}

	repoName, err := GetRepoName(Cfg.DocumentGitUrl)
	fmt.Println(Cfg.CurrentDir)
	fmt.Println(repoName)
	if err != nil {
		panic(err)
	}

	Cfg.AppName = "Boke"
	Cfg.Version = 3.0
	Cfg.DocumentDir = Cfg.CurrentDir + "/" + repoName
	Cfg.GitHookUrl = "/api/git_push_hook"
	Cfg.AppRepository = "git@github.com:Downsonwang/myblog"
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
	if !utils.IsDir(Cfg.DocumentDir) {
		fmt.Println("正在克隆文档仓库，请稍等...")
		out, err := utils.RunCmdByDir(Cfg.CurrentDir, "git", "clone", Cfg.DocumentGitUrl)
		fmt.Println(Cfg.CurrentDir)
		fmt.Println(Cfg.DocumentGitUrl)
		if err != nil {
			panic(err)
		}
		//fmt.Println(Cfg.CurrentDir)
		//fmt.Println(Cfg.DocumentGitUrl)
		fmt.Println(out)
	} else {
		out, err := utils.RunCmdByDir(Cfg.DocumentDir, "git", "pull")
		fmt.Println(Cfg.DocumentDir)
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
	if !utils.IsDir(imgDir) {
		if os.Mkdir(imgDir, os.ModePerm) != nil {
			panic("生成images目录失败！")
		}
	}
}

func checkDocDirAndBindConfig(cfg *Config) error {
	dirs := []string{"assets", "content", "extra_nav"}
	for _, dir := range dirs {
		absoluteDir := Cfg.DocumentDir + "/" + dir
		if !utils.IsDir(absoluteDir) {
			return errors.New("documents cannot lack " + absoluteDir + " dir")
		}
	}
	cfg.DocumentAssetsDir = cfg.DocumentDir + "/assets"
	cfg.DocumentContentDir = cfg.DocumentDir + "/content"
	cfg.DocumentExtraNavDir = cfg.DocumentDir + "/extra_nav"
	return nil
}
