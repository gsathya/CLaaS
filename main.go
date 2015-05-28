package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	queryString := r.FormValue("q")
	if len(queryString) > 0 {
		fmt.Fprintf(w, strings.ToUpper(queryString))
		return
	}

	fmt.Fprintf(w, "Try ?q=yolo")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
