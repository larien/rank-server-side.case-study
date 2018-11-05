package main

import (
	"fmt"

	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/routing"
)

const port = ":8080"

func main() {
	Rank()
}

// Rank starts the routine for Rank's app.
func Rank() {
	router := routing.Router()

	fmt.Println(router)

	router.Run(port)
}
