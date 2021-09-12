package utils

import (
	"os"
	"time"
)

const (
	backupMetaPath = "../test/backup.meta"
)

type RecoverInfo struct {
	FileType	string			`json:"file_type"`
	Mode		os.FileMode		`json:"mode"`
	UId			uint32			`json:"u_id"`
	GId			uint32			`json:"g_id"`
	ATime		time.Time		`json:"a_time"`
	MTime		time.Time		`json:"m_time"`
	SrcPath		string			`json:"src_path"`
	CopiedPath	string			`json:"copied_path"`
	LinkedPath	string			`json:"linked_path"`
}