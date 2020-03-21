package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// view 设置模板
func view(w http.ResponseWriter, r *http.Request) {
	//定义一个函数Parsize
	Prasize := func(name string) (string, error) {
		return name + "年轻又帅气", nil
	}
	//定义模板
	tem := template.New("view.tmpl") //创建一个模板对象
	//告诉模板，现在多了一个自定义函数
	tem.Funcs(template.FuncMap{
		"kua": Prasize,
	})

	_, err := tem.ParseFiles("./view.tmpl")
	//解析模板
	if err != nil {
		fmt.Println("Parse template faild err=", err)
		return
	}
	name := "卡布达"
	//渲染模板
	tem.Execute(w, name)
}

func demo01(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	tmpl, err := template.ParseFiles("./t.tmpl", "./ul.tmpl", "ol.tmpl")
	if err != nil {
		fmt.Println("Parse template faild err=", err)
		return
	}
	name := "奥特曼"
	//渲染模板
	tmpl.Execute(w, name)
}

func main() {
	http.HandleFunc("/", view)      //注册路由处理函数
	http.HandleFunc("/tem", demo01) //注册路由处理函数
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Http Server faild err=", err)
		return
	}

}
