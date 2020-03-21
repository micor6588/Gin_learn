package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	u1 := UserInfo{
		Name:   "奥特曼",
		Gender: "男",
		Age:    18,
	}

	u2 := UserInfo{
		Name:   "小怪兽",
		Gender: "男",
		Age:    26,
	}
	//将结构体写入到模板
	tmpl.Execute(w, map[string]interface{}{
		"u1": u1,
		"u2": u2,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9096", nil)
	if err != nil {
		fmt.Println("Http server faild,err=", err)
		return
	}

}
