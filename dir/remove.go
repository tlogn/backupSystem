package dir

import (
	"backupSystem/utils"
	"fmt"
	"log"
	"net/http"
	"os"
)

func LocalRemove(w http.ResponseWriter, r *utils.Request){
	srcPath := r.DirPara.DirPath
	fmt.Fprintf(w, "%v", localRemove(srcPath))
}

func localRemove(srcPath string) string{
	err := os.RemoveAll(srcPath)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	return utils.SucceedResponse()
}