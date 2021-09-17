package rpc_utils

import "fmt"

type Request struct {
	Username	string
	SrcPath		string
	CopiedPath	string
	ProcessPath	string
}

func (req Request) String()	string {
	return fmt.Sprintf(`{"Username":%s,"SrcPath":%s,"CopiedPath":%s,"ProcessPath":%s}`, req.Username, req.SrcPath, req.CopiedPath, req.ProcessPath)
}