package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type CpPara struct {
	OriginPath 	string	`json:"origin_path"`
	BackupPath	string	`json:"backup_path"`
}

type RecoverPara struct {
	RecoverPath string	`json:"recover_path"`
}

type  CompressPara struct {
	IsCompress		bool	`json:"is_compress"`
	CompressPath	string	`json:"compress_path"`
}

type EncodePara struct {
	IsEncode	bool	`json:"is_encode"`
	EncodePath	string	`json:"encode_path"`
}

type PackPara struct {
	isPack		bool	`json:"is_pack"`
	PackPath	string	`json:"pack_path"`
}

type Request struct {
	RemoteOp 		bool			`json:"remote_op"`

	CopyOp 			bool			`json:"copy_op"`
	CpPara			CpPara			`json:"cp_para"`

	RecoverOp		bool			`json:"recover_op"`
	RecoverPara 	RecoverPara		`json:"recover_para"`

	CompressOp		bool			`json:"compress_op"`
	CompressPara 	CompressPara	`json:"compress_para"`

	EncodeOp		bool			`json:"encode_op"`
	EncodePara		EncodePara		`json:"encode_para"`

	PackOp			bool			`json:"pack_op"`
	PackPara		PackPara		`json:"pack_para"`

}

func method(w http.ResponseWriter, r *http.Request) {
	// 跨域设置
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json") //返回数据格式是json

	r.ParseForm()  // 解析参数，默认是不会解析的
	log.Println(r.Form)  // 这些信息是输出到服务器端的打印信息
	log.Println("path", r.URL.Path)
	for k, v := range r.Form {
		log.Println("key:", k)
		log.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "%v", r.Form["body"]) // 这个写入到 w 的是输出到客户端的
}

func main() {
	http.HandleFunc("/method", method) // 设置访问的路由
	err := http.ListenAndServe(":8090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}