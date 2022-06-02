package router

import (
	"api/controller"
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequests() {
	http.HandleFunc("/posts", controller.GetAllArticles)

	http.ListenAndServe(":9876", nil)
}
