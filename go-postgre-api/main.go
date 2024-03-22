package main

import (
	"fmt"
	"go-postgre-api/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
