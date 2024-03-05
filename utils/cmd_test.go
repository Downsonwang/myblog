/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-12-17 17:34:16
 * @LastEditTime: 2024-03-04 19:12:07
 */
package utils_test

import (
	"blog/utils"
	"testing"
)

func TestRunCmdByDir(t *testing.T) {
	_, err := utils.RunCmdByDir("./", "ping", "127.0.0.1")
	if err != nil {
		t.Error("run cmd error", err)
	}
}
