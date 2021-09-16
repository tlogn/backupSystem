package login

import (
	"backupSystem/utils"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *utils.Request){
	fmt.Fprintf(w, "%v", login(r.LoginPara.Username, r.LoginPara.Password))
}

func login(username string, password string) string {
	value, err := utils.RedisClient.Get(utils.Ctx, username).Result()
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	if value == "" {
		err = errors.New("username not exists")
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	if value != password {
		err = errors.New("wrong password")
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	return utils.SucceedResponse()
}