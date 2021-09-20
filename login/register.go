package login

import (
	"backupSystem/rpc/client"
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
	if utils.IsRedisKeyExist(username) {
		err := errors.New("username already exists")
		log.Println(err)
		return utils.ErrorResponse(err)
	}

	err := utils.RedisClient.Set(utils.Ctx, username, password, 0).Err()
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	client.RemoteMkdir("/home/lighthouse/backup/" + username)
	return utils.SucceedResponse()
}