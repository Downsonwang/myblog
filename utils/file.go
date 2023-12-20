/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-12-19 20:32:54
 * @LastEditTime: 2023-12-19 20:33:06
 */
package utils

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func IsDir(name string) bool {
	if info, err := os.Stat(name); err == nil {
		return info.IsDir()
	}
	return false
}

func IsFile(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func MakeDir(dir string) error {
	if !IsDir(dir) {
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}

func RemoveDir(dir string) error {
	if !IsDir(dir) {
		return errors.New("Cannot delete without directroy")
	}
	return os.RemoveAll(dir)
}

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	des, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer des.Close()
	nBytes, err := io.Copy(des, source)
	return nBytes, err
}
