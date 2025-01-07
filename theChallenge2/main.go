package main

import (
	"fmt"

	"math/rand"
	"time"

	"theChallenge2/controllers"
	"theChallenge2/routes"

	"github.com/starkbank/sdk-go/starkbank"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	starkbank.User = controllers.Auth()
	fmt.Println("Iniciando servido Rest com GO")
	routes.HandleRequest()
}
