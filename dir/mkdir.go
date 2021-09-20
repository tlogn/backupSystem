package dir

import (
	"backupSystem/utils"
	"fmt"
	"log"
	"net/http"
	"os"
)

func LocalMkdir(w http.ResponseWriter, r *utils.Request){
	srcPath := r.DirPara.DirPath
	fmt.Fprintf(w, "%v", localMkdir(srcPath))
}

func localMkdir(srcPath string) string{
	err := os.MkdirAll(srcPath, 0777)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	return utils.SucceedResponse()
}