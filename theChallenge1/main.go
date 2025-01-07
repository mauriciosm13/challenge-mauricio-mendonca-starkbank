package main

import (
	"math/rand"
	"time"

	"challenge-mauricio-mendonca-starkbank/theChallenge1/controllers"

	"github.com/starkbank/sdk-go/starkbank"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	starkbank.User = controllers.Auth()

	controllers.GenerateInvoiceAwaite()

}
