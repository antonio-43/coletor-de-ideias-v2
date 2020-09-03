package main

import (
    "fmt"
    "cdi/router"
    "net/http"
    "log"
)

func main() {
    r := router.Router()

    fmt.Println("API - fist golang api")
    fmt.Println("#GooglePlzHireMe")

    log.Fatal(http.ListenAndServe(":8080", r))
}
