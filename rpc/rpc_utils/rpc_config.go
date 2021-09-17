package rpc_utils

type Request struct {
	Username	string	`json:"username"`
	SrcPath		string	`json:"src_path"`
	CopiedPath	string	`json:"copied_path"`
	ProcessPath	string	`json:"process_path"`
}