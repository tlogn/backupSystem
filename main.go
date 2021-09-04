package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  // 解析参数，默认是不会解析的
	fmt.Println(r.Form)  // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	//s, _ := os.Stat("./copy")
	//fmt.Fprintf(w, "%v", s.Sys()) // 这个写入到 w 的是输出到客户端的
}

func pr(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json") //返回数据格式是json

	r.ParseForm()  // 解析参数，默认是不会解析的
	fmt.Println(r.Form)  // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	s, _ := os.Stat("./copy")
	fmt.Fprintf(w, "我操你妈的傻逼跨域%v", s.Sys()) // 这个写入到 w 的是输出到客户端的
}

func main() {
	http.HandleFunc("/method", pr) // 设置访问的路由
	http.HandleFunc("/", sayhelloName) // 设置访问的路由
	err := http.ListenAndServe(":8090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}