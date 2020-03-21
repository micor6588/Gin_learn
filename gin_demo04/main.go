package main

import (
	"fmt"
	"net/http"
)

///index 设置模板
func index(w http.ResponseWriter, r *http.Request) {
	//定义模板

	//解析模板

	//渲染模板
}

//
func home(w http.ResponseWriter, r *http.Request) {
	//定义模板

	//解析模板

	//渲染模板
}
func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Http Server faild err=", err)
		return
	}

}
