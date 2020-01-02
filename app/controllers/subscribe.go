package controllers

import (
	"github.com/residenti/trading_bitcoin_api/app/models"
	"github.com/residenti/trading_bitcoin_api/bitflyer"
	"github.com/residenti/trading_bitcoin_api/config"
)

func SubscribeData() {
	apiClinet := bitflyer.New()
	tickerChannel := make(chan bitflyer.Ticker)
	go apiClinet.SubscribeTicker(config.List.ProductCode, tickerChannel)

	go func() {
		for ticker := range tickerChannel {
			for _, duration := range config.List.Durations {
				isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
				if isCreated == true && duration == config.List.TradeDuration {
					// TODO
				}
			}
		}
	}()
}
