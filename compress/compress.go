package compress

import (
	"backupSystem/utils"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Compress(w http.ResponseWriter, r *utils.Request){
	comOrUncom := r.CompressPara.IsCompress
	if comOrUncom {
		fmt.Fprintf(w, "%v", compress(r))
	} else {
		fmt.Fprintf(w, "%v", undoCompress(r))
	}
}

func build


func compress(r *utils.Request) error {

	srcPath, err := filepath.Abs(r.CompressPara.CompressPath)

	if err != nil {
		log.Println(err)
		return err
	}

	if utils.IsDir(srcPath) {
		err := errors.New("is a fucking dir")
		log.Printf("%v %v", srcPath, err)
		return err
	}

	file, err := ioutil.ReadFile(srcPath)
	if err != nil {
		log.Printf("read file %v error, %v",srcPath, err)
		return err
	}

	err = ioutil.WriteFile(srcPath + ".ylxcompress", compressedFile, 0777)
	return nil
}

func undoCompress(r *utils.Request) error {

	return nil
}

