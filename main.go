package main

import (
	"cdi/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	fmt.Println("API - fist golang api")
	fmt.Println("#GooglePlzHireMe")

	log.Fatal(http.ListenAndServe(":8080", r))
}
