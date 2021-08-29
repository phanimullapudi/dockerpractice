package main

import (
	"fmt"

	"github.com/phanimullapudi/dockerpractice/E2EApplication/internal/database"
)

type Poll struct {
	Name   string `json:"name"`
	Choice int    `json:"age"`
}

func main() {

	redisclient := database.NewRedisConnect()

	pong, err := redisclient.Ping().Result()
	fmt.Println(pong, err)

}
