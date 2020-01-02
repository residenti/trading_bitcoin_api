package main

import (
	"github.com/residenti/trading_bitcoin_api/app/controllers"
	"github.com/residenti/trading_bitcoin_api/config"
	"github.com/residenti/trading_bitcoin_api/utils"
)

func init() {
	utils.InitSettingsOfLog(config.List.Logfile)
}

func main() {
	controllers.SubscribeData()
	controllers.StartServer()
}
