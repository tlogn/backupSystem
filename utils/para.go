package utils

type DirPara struct {
	DirPath string	`json:"dir_path"`
}

type CopyPara struct {
	OriginPath 	string	`json:"origin_path"`
	BackupPath	string	`json:"backup_path"`
}

type CompressPara struct {
	IsCompress		bool	`json:"is_compress"`
	CompressPath	string	`json:"compress_path"`
}

type EncodePara struct {
	IsEncode	bool	`json:"is_encode"`
	EncodePath	string	`json:"encode_path"`
}

type PackPara struct {
	IsPack		bool	`json:"is_pack"`
	PackPath	string	`json:"pack_path"`
}

type RecoverPara struct {
	RecoverPath string	`json:"recover_path"`
}

type LoginPara struct {
	Username string	`json:"username"`
	Password string	`json:"password"`
}