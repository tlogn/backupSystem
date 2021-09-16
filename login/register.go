package login

import (
	"backupSystem/utils"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *utils.Request){
	fmt.Fprintf(w, "%v",register(r.LoginPara.Username, r.LoginPara.Password))
}

func register(username string, password string) string {
	value, err := utils.RedisClient.Get(utils.Ctx, username).Result()
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	if value == "" {
		err = utils.RedisClient.Set(utils.Ctx, username, password, 0).Err()
		if err != nil {
			log.Println(err)
			return utils.ErrorResponse(err)
		}
		return utils.SucceedResponse()
	}
	err = errors.New("username already exists")
	log.Println(err)
	return utils.ErrorResponse(err)
}