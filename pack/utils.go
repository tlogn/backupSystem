package pack

import (
	"backupSystem/utils"
	"log"
)

func LPack(packPath string) error {
	errResp := localPack(packPath)
	if errResp != utils.SucceedResponse() {
		err := utils.GetErrFromResponse(errResp)
		log.Println(err)
		return err
	}
	return nil
}

func UPack(packPath string) error {
	errResp := localUnpack(packPath)
	if errResp != utils.SucceedResponse() {
		err := utils.GetErrFromResponse(errResp)
		log.Println(err)
		return err
	}
	return nil
}