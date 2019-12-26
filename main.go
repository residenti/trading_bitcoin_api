package main

import (
	"fmt"
	"net/http"

	"rsc.io/quote"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, quote.Hello())
}
