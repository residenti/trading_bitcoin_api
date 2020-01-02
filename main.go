package main

import (
	"fmt"
	"net/http"

	"github.com/residenti/trading_bitcoin_api/app/controllers"
	"github.com/residenti/trading_bitcoin_api/config"
	"github.com/residenti/trading_bitcoin_api/utils"
)

func init() {
	utils.InitSettingsOfLog(config.List.Logfile)
}

func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)

	controllers.SubscribeData()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
