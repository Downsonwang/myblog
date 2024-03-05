/*
 * @Descripttion:
 * @Author: DW
 * @Date: 2023-12-17 17:34:16
 * @LastEditTime: 2024-03-04 19:15:47
 */
package controller

import (
	"blogdemo/models"
	"fmt"
	"net/http"
)

func GithubHook(w http.ResponseWriter, r *http.Request) {

	SedResponse(w, "ok")

	models.CompiledContent()
}

func SedResponse(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write([]byte(`{"msg": "` + msg + `"}`))
	if err != nil {
		fmt.Println(err)
	}
}
