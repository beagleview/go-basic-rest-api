package handles

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Newbie and Noob")
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage).Methods("GET")
	return r
}
