package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	FILE_TYPE_DIR 		= "Dir"
	FILE_TYPE_FILE		= "File"
	FILE_TYPE_PIPELINE 	= "Pipeline"
	FILE_TYPE_SYMLINK 	= "Symlink"
	FILE_TYPE_HARDLINK	= "Hardlink"
)

type RecoverInfo struct {
	FileType	string			`json:"file_type"`
	Mode		os.FileMode		`json:"mode"`
	UId			int				`json:"u_id"`
	GId			int				`json:"g_id"`
	ATime		time.Time		`json:"a_time"`
	MTime		time.Time		`json:"m_time"`
	SrcPath		string			`json:"src_path"`
	CopiedPath	string			`json:"copied_path"`
	LinkedPath	string			`json:"linked_path"`
	DirList		[]string		`json:"dir_list"`
}

type Request struct {
	Op 				string			`json:"op"`

	UserName		string			`json:"user_name"`

	TransPara		TransPara		`json:"trans_para"`

	LoginPara		LoginPara		`json:"login_para"`

	DirPara			DirPara			`json:"dir_para"`

	CopyPara		CopyPara		`json:"copy_para"`

	RecoverPara 	RecoverPara		`json:"recover_para"`

	CompressPara 	CompressPara	`json:"compress_para"`

	EncodePara		EncodePara		`json:"encode_para"`

	PackPara		PackPara		`json:"pack_para"`

}

func (request *Request) SetRequestFromHttp(r *http.Request) error {
	if r.Method == "GET" {
		err := json.Unmarshal([]byte(r.Form.Get("body")), &request)
		if err != nil {
			return err
		}
	} else if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(body, &request)
		if err != nil {
			return err
		}
	}
	return nil
}

type DirFile struct {
	FileName 	string	`json:"file_name"`
	IsDir		bool	`json:"is_dir"`
	FilePath	string	`json:"file_path"`
}

type Response struct {
	Succeed		bool		`json:"succeed"`
	Err			string 		`json:"err"`
	DirFiles	[]DirFile 	`json:"dir_files"`
	FileType	string		`json:"file_type"`
	Data		[]byte		`json:"data"`
}