package main

import (
	"fmt"
	"net/http"

	"github.com/residenti/trading_bitcoin_api/config"

	"rsc.io/quote"
)

func main() {
	fmt.Println(config.List.Item)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, quote.Hello())
}
