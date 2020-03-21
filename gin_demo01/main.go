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
	user := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, user)
}

func main(){
	http.HandleFunc("/",sayHello)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil{
	fmt.Println("Http server faild,err=",err)
		return
	}

}
