package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// UserInfo 定义用户结构体
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

	hobbyList := []string{
		"篮球",
		"足球",
		"乒乓球",
		"羽毛球",
		"口风琴",
		"小提琴",
	}
	//将结构体写入到模板
	tmpl.Execute(w, map[string]interface{}{
		"u1":    u1,
		"u2":    u2,
		"hobby": hobbyList,
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
