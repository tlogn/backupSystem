package main

import (
	"backupSystem/compress"
	"backupSystem/copy"
	"backupSystem/dir"
	"backupSystem/encode"
	"backupSystem/login"
	"backupSystem/pack"
	"backupSystem/recover"
	"backupSystem/rpc/client"
	"backupSystem/rpc/server"
	"backupSystem/utils"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type HTTPHandler func(w http.ResponseWriter, r *utils.Request)

var (
	route = map[string]HTTPHandler {
		"login": login.Login,
		"register": login.Register,
		"local_copy" : copy.LocalCpFile,
		"local_filter_copy" : copy.LocalFilterCpFile,
		"local_dir" : dir.LocalDir,
		"local_encode" : encode.LocalEncode,
		"local_compress" : compress.LocalCompress,
		"local_recover" : recover.LocalRecover,
		"local_pack" : pack.LocalPack,
		"local_remove" : dir.LocalRemove,
		"local_mkdir" : dir.LocalMkdir,
		"local_cmp"	: dir.LocalCmp,
		"remote_dir" : client.RemoteDir,
		"remote_encode" : client.RemoteEncode,
		"remote_compress" : client.RemoteCompress,
		"remote_pack" : client.RemotePack,
		"remote_download" : client.RemoteDownload,
		"remote_upload" : client.RemoteUpload,
		"remote_remove" : client.RemoteRemove,
		"remote_mkdir" : client.RemoteMkDir,
	}
	rpcServerSelect bool
)

func method(w http.ResponseWriter, r *http.Request) {
	// 跨域设置
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json") //返回数据格式是json

	r.ParseForm()  // 解析参数，默认是不会解析的

	request := utils.Request{}

	err := request.SetRequestFromHttp(r)
	fmt.Println(request, request.Op)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%v", utils.ErrorResponse(err))
		return
	}

	handler, ok:= route[request.Op]
	if !ok {
		err := errors.New("illegal route")
		log.Println(err)
		fmt.Fprintf(w, "%v", utils.ErrorResponse(err))
		return
	}

	handler(w, &request)
}

func main() {
	flag.BoolVar(&rpcServerSelect, "server", false, "server option")
	flag.Parse()
	if rpcServerSelect {
		server.RunRpcServer()
	} else {
		http.HandleFunc("/method", method)       // 设置访问的路由
		err := http.ListenAndServe(":8090", nil) // 设置监听的端口
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}
}

/*
{
	"op":"remote_upload",

	"user_name":"",

	"filter_path":"",

	"trans_para":{
		"local_path":"/Users/bytedance/go/src/backupSystem/compress/test",
		"remote_path":"/home/lighthouse/go/src/test"
	},

	"login_para":{
		"username":"",
		"password":""
	},

	"dir_para":{
		"dir_path":""
	},

	"copy_para":{
		"origin_path":"",
		"backup_path":""
	},

	"recover_para":{
		"recover_path":""
	},

	"compress_para":{
		"is_compress":false,
		"compress_path":""
	},

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