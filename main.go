package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/residenti/trading_bitcoin_api/bitflyer"
	"github.com/residenti/trading_bitcoin_api/config"
	"github.com/residenti/trading_bitcoin_api/utils"
)

func init() {
	utils.InitSettingsOfLog(config.List.Logfile)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	apiClinet := bitflyer.New()
	ticker, err := apiClinet.GetTicker("BTC_JPY")
	if err != nil {
		log.Printf("handler err=%s", err.Error())
	}

	fmt.Fprint(w, *ticker)
}
