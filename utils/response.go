package utils

type DirFile struct {
	FileName 	string	`json:"file_name"`
	IsDir		bool	`json:"is_dir"`
	FilePath	string	`json:"file_path"`
}

type Response struct {
	Succeed		bool		`json:"succeed"`
	Err			string 		`json:"err"`
	DirFiles	[]DirFile 	`json:"dir_files"`
}