package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/residenti/trading_bitcoin_api/config"
	"github.com/residenti/trading_bitcoin_api/utils"

	"rsc.io/quote"
)

func init() {
	utils.InitSettingsOfLog(config.List.Logfile)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("called!")
	fmt.Fprintf(w, quote.Hello())
}
