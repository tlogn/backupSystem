package main

import (
	"backupSystem/compress"
	"backupSystem/copy"
	"backupSystem/dir"
	"backupSystem/encode"
	"backupSystem/pack"
	"backupSystem/recover"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Request struct {
	RemoteOp 		bool			`json:"remote_op"`

	GetDirOp		bool			`json:"get_dir_op"`
	GetDirPara		dir.Para		`json:"get_dir_para"`

	CopyOp 			bool			`json:"copy_op"`
	CopyPara		copy.Para		`json:"copy_para"`

	RecoverOp		bool			`json:"recover_op"`
	RecoverPara 	recover.Para	`json:"recover_para"`

	CompressOp		bool			`json:"compress_op"`
	CompressPara 	compress.Para	`json:"compress_para"`

	EncodeOp		bool			`json:"encode_op"`
	EncodePara		encode.Para		`json:"encode_para"`

	PackOp			bool			`json:"pack_op"`
	PackPara		pack.Para		`json:"pack_para"`

}

type DirFile struct {
	FileName 	string	`json:"file_name"`
	IsDir		bool	`json:"is_dir"`
}

type Response struct {
	Succeed		bool		`json:"succeed"`
	Err			string		`json:"err"`
	DirFiles	[]DirFile	`json:"dir_files"`
}

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