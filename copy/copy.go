package copy

import (
	"backupSystem/utils"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"syscall"
	"time"
)

func CpFile(dstPath, srcPath string) error {

	return nil
}

func CpNormalFile(dstPath, srcPath string) error {
	f, err:= os.Stat(srcPath)
	if err != nil {
		log.Fatalf("open file %v error, %v",srcPath, err)
		return err
	}

	if f.IsDir() {
		err := os.MkdirAll(dstPath, 0644)
		if err != nil {
			log.Fatalf("mkdir error, %v", err)
			return err
		}
		recoverInfo := utils.RecoverInfo{
			FileType: "Dir",
			Mode: f.Mode(),
			UId: f.Sys().(*syscall.Stat_t).Uid,
			GId: f.Sys().(*syscall.Stat_t).Gid,
			ATime: time.Unix(f.Sys().(*syscall.Stat_t).Atimespec.Sec, f.Sys().(*syscall.Stat_t).Atimespec.Nsec),
			MTime: time.Unix(f.Sys().(*syscall.Stat_t).Mtimespec.Sec, f.Sys().(*syscall.Stat_t).Mtimespec.Nsec),
			SrcPath: srcPath,
			CopiedPath: dstPath,
		}
		recoverInfoJson, _ := json.Marshal(recoverInfo)
		utils.RedisClient.Set(utils.Ctx, "local_"+dstPath, recoverInfoJson, 0)
		return nil
	}

	file, err := ioutil.ReadFile(srcPath)
	if err != nil {
		log.Fatalf("read file %v error, %v",srcPath, err)
		return err
	}
	err = ioutil.WriteFile(dstPath, file, 0644)
	if err != nil {
		log.Fatalf("write file %v error, %v",dstPath, err)
		return errors.New("write file error")
	}
	recoverInfo := utils.RecoverInfo{
		FileType: "File",
		Mode: f.Mode(),
		UId: f.Sys().(*syscall.Stat_t).Uid,
		GId: f.Sys().(*syscall.Stat_t).Gid,
		ATime: time.Unix(f.Sys().(*syscall.Stat_t).Atimespec.Sec, f.Sys().(*syscall.Stat_t).Atimespec.Nsec),
		MTime: time.Unix(f.Sys().(*syscall.Stat_t).Mtimespec.Sec, f.Sys().(*syscall.Stat_t).Mtimespec.Nsec),
		SrcPath: srcPath,
		CopiedPath: dstPath,
	}
	recoverInfoJson, _ := json.Marshal(recoverInfo)
	utils.RedisClient.Set(utils.Ctx, "local_"+dstPath, recoverInfoJson, 0)
	return nil
}

func CpHardLink(dstPath, srcPath string) error {
	err := os.Link(srcPath, dstPath)
	if err != nil {
		log.Println(err)
		return err
	}

	f, _ := os.Lstat(srcPath)
	recoverInfo := utils.RecoverInfo{
		FileType: "HardLink",
		Mode: f.Mode(),
		UId: f.Sys().(*syscall.Stat_t).Uid,
		GId: f.Sys().(*syscall.Stat_t).Gid,
		ATime: time.Unix(f.Sys().(*syscall.Stat_t).Atimespec.Sec, f.Sys().(*syscall.Stat_t).Atimespec.Nsec),
		MTime: time.Unix(f.Sys().(*syscall.Stat_t).Mtimespec.Sec, f.Sys().(*syscall.Stat_t).Mtimespec.Nsec),
		SrcPath: srcPath,
		CopiedPath: dstPath,
	}
	recoverInfoJson, _ := json.Marshal(recoverInfo)
	utils.RedisClient.Set(utils.Ctx, "local_"+dstPath, recoverInfoJson, 0)
	return nil
}

func CpSymLink(dstPath, srcPath string) error {
	f, _ := os.Lstat(srcPath)
	err := CpNormalFile(dstPath, srcPath)
	if err != nil {
		return err
	}

	linkedSrcPath, _ := os.Readlink(srcPath)
	absLinkedSrcPath, _ := filepath.Abs(linkedSrcPath)
	recoverInfo := utils.RecoverInfo{
		FileType: "SymLink",
		Mode: f.Mode(),
		UId: f.Sys().(*syscall.Stat_t).Uid,
		GId: f.Sys().(*syscall.Stat_t).Gid,
		ATime: time.Unix(f.Sys().(*syscall.Stat_t).Atimespec.Sec, f.Sys().(*syscall.Stat_t).Atimespec.Nsec),
		MTime: time.Unix(f.Sys().(*syscall.Stat_t).Mtimespec.Sec, f.Sys().(*syscall.Stat_t).Mtimespec.Nsec),
		SrcPath: srcPath,
		CopiedPath: dstPath,
		LinkedPath: absLinkedSrcPath,
	}
	recoverInfoJson, _ := json.Marshal(recoverInfo)
	utils.RedisClient.Set(utils.Ctx, "local_"+dstPath, recoverInfoJson, 0)
	return nil
}

func CpPipeline(dstPath, srcPath string) error {
	err := syscall.Mkfifo(dstPath, 0644)
	if err != nil {
		return err
	}
	f, _ := os.Lstat(srcPath)
	recoverInfo := utils.RecoverInfo{
		FileType: "Pipeline",
		Mode: f.Mode(),
		UId: f.Sys().(*syscall.Stat_t).Uid,
		GId: f.Sys().(*syscall.Stat_t).Gid,
		ATime: time.Unix(f.Sys().(*syscall.Stat_t).Atimespec.Sec, f.Sys().(*syscall.Stat_t).Atimespec.Nsec),
		MTime: time.Unix(f.Sys().(*syscall.Stat_t).Mtimespec.Sec, f.Sys().(*syscall.Stat_t).Mtimespec.Nsec),
		SrcPath: srcPath,
		CopiedPath: dstPath,
	}
	recoverInfoJson, _ := json.Marshal(recoverInfo)
	utils.RedisClient.Set(utils.Ctx, "local_"+dstPath, recoverInfoJson, 0)
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