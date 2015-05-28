// To the extent possible under law, Sathyanarayanan Gunasekaran waived
// all copyright and related or neighboring rights to CLaaS, using the
// creative commons "cc0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

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
