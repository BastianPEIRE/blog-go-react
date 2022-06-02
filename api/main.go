package main

import (
	"api/router"
	_ "github.com/lib/pq"
)

func main() {
	router.HandleRequests()
}
