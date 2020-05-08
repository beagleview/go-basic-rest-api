package handles

import (
	"encoding/json"
	"fmt"
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
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}
