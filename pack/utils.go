package pack

import (
	"backupSystem/utils"
	"errors"
	"log"
)

func LPack(packPath string) error {
	if localPack(packPath) != utils.SucceedResponse() {
		err := errors.New("pack error")
		log.Println(err)
		return err
	}
	return nil
}

func UPack(packPath string) error {
	if localUnpack(packPath) != utils.SucceedResponse() {
		err := errors.New("unpack error")
		log.Println(err)
		return err
	}
	return nil
}