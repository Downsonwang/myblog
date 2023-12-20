/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-12-19 20:25:19
 * @LastEditTime: 2023-12-19 20:25:52
 */
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