package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Test Content 1"},
		Article{Title: "Newbie Title", Desc: "Newbie Description", Content: "Newbie and Noob Content"},
	}

	fmt.Println("Endpoint Hit : Get all articles")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Newbie and Noob")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", getAllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
