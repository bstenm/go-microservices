package main

import (
	"fmt"

	"github.com/callistaenterprise/goblog/quoteservice/service"
)

var appName = "quoteService"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6768")
}
