package controller

import (
	"api/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type Posts struct {
	Id      int    `json:"id"`
	Img     string `json:"img"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
}

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("1")
	db := service.DbConnect()
	var (
		id          int
		img         string
		title       string
		content     string
		userId      int
		finalResult []Posts
	)

	result, err := db.Query(`SELECT * FROM "schema-blog".posts`)
	if err != nil {
		fmt.Println(err)
	} else {
		for result.Next() {
			error := result.Scan(&id, &img, &title, &content, &userId)
			if error != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%v, %v, %v, %v, %v\n", id, img, title, content, userId)
			finalResult = append(finalResult, Posts{id, img, title, content, userId})
		}
		err := json.NewEncoder(w).Encode(finalResult)
		if err != nil {
			return
		}
	}
}
