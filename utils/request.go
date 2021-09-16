package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Op 				string			`json:"op"`

	UserName		string			`json:"user_name"`

	LoginPara		LoginPara		`json:"login_para"`

	GetDirPara		DirPara			`json:"get_dir_para"`

	CopyPara		CopyPara		`json:"copy_para"`

	RecoverPara 	RecoverPara		`json:"recover_para"`

	CompressPara 	CompressPara	`json:"compress_para"`

	EncodePara		EncodePara		`json:"encode_para"`

	PackPara		PackPara		`json:"pack_para"`

}

func (request *Request) SetRequest(r *http.Request) error {
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