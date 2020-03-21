package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//方法一：
		// data := map[string]interface{}{
		// 	"name":    "刘惠玲",
		// 	"message": "hello World",
		// 	"age ":    26,
		// }
		data := gin.H{
			"name":    "刘惠玲",
			"message": "hello World",
			"age ":    20,
		}
		c.JSON(http.StatusOK, data)

	})
	//方法2;灵活使用Tag来对结构体字段做定制化操作
	type msg struct {
		Name    string
		Message string
		Age     int
	}
	r.GET("/msg_json", func(c *gin.Context) {
		data1 := msg{
			Name:    "关倩",
			Message: "打篮球",
			Age:     26,
		}
		c.JSON(http.StatusOK, data1) //本质还是json的序列化，首字母小写，不可导出

	})
	type user struct {
		name    string `json:"name"`
		message string
		age     int
	}
	r.GET("/user", func(c *gin.Context) {
		data2 := user{
			name:    "关倩",
			message: "打篮球",
			age:     26,
		}
		c.JSON(http.StatusOK, data2)
	}) //本质还是json的序列化，首字母小写，不可导出
	r.Run(":8085")
}
