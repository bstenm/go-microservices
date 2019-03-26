package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/callistaenterprise/goblog/accountservice/dbclient"
	"github.com/callistaenterprise/goblog/accountservice/service"
)

var appName = "accountService"

type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

// HomeHandler to handle request to home path
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Voila")
}

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
