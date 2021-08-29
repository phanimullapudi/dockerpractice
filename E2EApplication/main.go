package main

import (
	"fmt"

	redisConnectRepo "github.com/phanimullapudi/dockerpractice/E2EApplication/internal/database"
)

type Poll struct {
	Name   string `json:"name"`
	Choice int    `json:"age"`
}

func main() {

	redisclient := redisConnectRepo.ConnectDB()

	pong, err := redisclient.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

}
