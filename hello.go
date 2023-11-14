package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		variables := mux.Vars(r)
		title := variables["title"]
		page := variables["page"]

		fmt.Fprintf(w, "You have requested the book: %s on page %s\n", title, page)

	})

	http.ListenAndServe(":80", router)
}
