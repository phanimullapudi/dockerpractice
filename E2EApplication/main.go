package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	redisConnectRepo "github.com/phanimullapudi/dockerpractice/E2EApplication/internal/database"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page")
	fmt.Println("Endpoint Hit: homePage")

	redisclient := redisConnectRepo.ConnectDB()
	err := redisclient.Set("name", "Elliot", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

}

func returnAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllItems")

	redisclient := redisConnectRepo.ConnectDB()
	val, err := redisclient.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(w, val)

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllItems)

	log.Fatal(http.ListenAndServe(":1000", myRouter))
}

func main() {

	redisclient := redisConnectRepo.ConnectDB()

	pong, err := redisclient.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)
	handleRequests()
}
