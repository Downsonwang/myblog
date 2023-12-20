/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-12-19 20:36:48
 * @LastEditTime: 2023-12-19 20:36:54
 */
package utils

import (
	"os/exec"
)

func RunCmdByDir(dir string, cmdName string, arg ...string) (string, error) {
	cmd := exec.Command(cmdName, arg...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
