package copy

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func CpFile(dstPath, srcPath string) error {
	f, err := ioutil.ReadFile(srcPath)
	if err != nil {
		log.Fatalf("read file %v error, %v",srcPath, err)
		return err
	}
	err = ioutil.WriteFile(dstPath, f, 0644)
	if err != nil {
		log.Fatalf("write file %v error, %v",dstPath, err)
		return errors.New("write file error")
	}
	return nil
}

func CpDir(dstPath, srcPath string) error {
	return cpDir(path.Join(dstPath, path.Base(srcPath)), srcPath)
}

func cpDir(dstPath, srcPath string) error {

	fileList, err := ioutil.ReadDir(srcPath)
	if err != nil {
		log.Fatalf("read dir %v error, %v",srcPath, err)
		return err
	}

	err = os.MkdirAll(dstPath, 0777)
	if err != nil {
		log.Fatalf("mkdir %v error, %v", path.Join(srcPath, path.Base(dstPath)), err)
		return err
	}

	for _, f := range fileList {
		if f.IsDir() {
			cpDir(path.Join(dstPath, f.Name()), path.Join(srcPath, f.Name()))
		} else {
			CpFile(path.Join(dstPath, f.Name()), path.Join(srcPath, f.Name()))
		}
	}
	return nil
}