package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func method(w http.ResponseWriter, r *http.Request) {
	// 跨域设置
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json") //返回数据格式是json

	r.ParseForm()  // 解析参数，默认是不会解析的
	log.Println(r.Form)  // 这些信息是输出到服务器端的打印信息
	log.Println("path", r.URL.Path)
	log.Println("scheme", r.URL.Scheme)
	log.Println(r.Form["url_long"])
	log.Println(r.Form["body"])
	for k, v := range r.Form {
		log.Println("key:", k)
		log.Println("val:", strings.Join(v, ""))
	}
	s, _ := os.Stat("./copy")
	fmt.Fprintf(w, "%v", s.Sys()) // 这个写入到 w 的是输出到客户端的
}

func main() {
	http.HandleFunc("/method", method) // 设置访问的路由
	err := http.ListenAndServe(":8090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}