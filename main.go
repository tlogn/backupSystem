package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type handler func(w http.ResponseWriter, r *http.Request)

var (
	route map[string]handler
)

func method(w http.ResponseWriter, r *http.Request) {
	// 跨域设置
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json") //返回数据格式是json

	r.ParseForm()  // 解析参数，默认是不会解析的

	log.Println("Form", r.Form)  // 这些信息是输出到服务器端的打印信息
	log.Println("Body", r.Form)  // 这些信息是输出到服务器端的打印信息
	log.Println("requ", r)

	for k, v := range r.Form {
		log.Println("key:", k)
		log.Println("val:", strings.Join(v, ""))
	}

	var a Request
	_ = json.Unmarshal([]byte(r.Form.Get("body")), a)
	fmt.Fprintf(w, "%v", r) // 这个写入到 w 的是输出到客户端的
}

func main() {

	http.HandleFunc("/method", method) // 设置访问的路由
	err := http.ListenAndServe(":8090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


	//var a Response
	//a.DirFiles = []DirFile{{"", false}, {"", false}}
	//b, _:= json.Marshal(a)
	//fmt.Println(string(b))
}

/*
{
	"remote_op":false,

	"get_dir_op":false,
	"get_dir_para":{
		"dir_path":""
	},

	"copy_op":false,
	"copy_para":{
		"origin_path":"",
		"backup_path":""
	},

	"recover_op":false,
	"recover_para":{
		"recover_path":""
	},

	"compress_op":false,
	"compress_para":{
		"is_compress":false,
		"compress_path":""
	},

	"encode_op":false,
	"encode_para":{
		"is_encode":false,
		"encode_path":""
	},

	"pack_para":{
		"is_pack":false,
		"pack_path":""
	}
}

{
	"succeed":false,
	"err":"",
	"dir_files":[
		{
			"file_name":"",
			"is_dir":false
		},
		{
			"file_name":"",
			"is_dir":false
		}
	]
}


*/