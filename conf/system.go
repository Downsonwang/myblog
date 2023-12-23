/*
 * @Descripttion: set the system info like currentdi and so on to get docs
 * @Author:DW
 * @Date: 2023-12-19 20:25:19
 * @LastEditTime: 2023-12-23 17:30:38
 */
// DW
package conf

type SystemInfo struct {
	AppName             string
	Version             float32
	CurrentDir          string
	GitHookUrl          string
	AppRepository       string
	DocumentDir         string
	DocumentAssetsDir   string
	DocumentContentDir  string
	DocumentExtraNavDir string
}
