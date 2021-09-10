package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HTTPHandler func(w http.ResponseWriter, r *Request)

var (
	route map[string]HTTPHandler = map[string]HTTPHandler {
		"local_copy" : func(w http.ResponseWriter, r *Request){},
		"local_dir" : func(w http.ResponseWriter, r *Request){},
		"local_encode" : func(w http.ResponseWriter, r *Request){},
		"local_compress" : func(w http.ResponseWriter, r *Request){},
		"local_recover" : func(w http.ResponseWriter, r *Request){},
		"local_pack" : func(w http.ResponseWriter, r *Request){},
		"remote_copy" : func(w http.ResponseWriter, r *Request){},
		"remote_dir" : func(w http.ResponseWriter, r *Request){},
		"remote_encode" : func(w http.ResponseWriter, r *Request){},
		"remote_compress" : func(w http.ResponseWriter, r *Request){},
		"remote_recover" : func(w http.ResponseWriter, r *Request){},
		"remote_pack" : func(w http.ResponseWriter, r *Request){},

	}
)

func method(w http.ResponseWriter, r *http.Request) {
	// 跨域设置
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json") //返回数据格式是json

	r.ParseForm()  // 解析参数，默认是不会解析的

	log.Println("Form", r.Form["body"])  // 这些信息是输出到服务器端的打印信息

	request := Request{}
	response := Response{Succeed: true}

	err := json.Unmarshal([]byte(r.Form.Get("body")), &request)
	if err != nil {
		log.Println(err)

		response.Succeed = false
		response.Err = err.Error()
		resp, _ := json.Marshal(response)

		fmt.Fprintf(w, "%v", string(resp))
		return
	}

	resp, _ := json.Marshal(response)
	fmt.Fprintf(w, "%v", string(resp)) // 这个写入到 w 的是输出到客户端的
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