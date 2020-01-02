package models

import (
	"fmt"
	"log"
	"time"

	"github.com/residenti/trading_bitcoin_api/bitflyer"
)

type Candle struct {
	ProductCode string        `json:"product_code"`
	Duration    time.Duration `json:"duration"`
	Time        time.Time     `json:"time"`
	Open        float64       `json:"open"`
	Close       float64       `json:"close"`
	High        float64       `json:"high"`
	Low         float64       `json:"low"`
	Volume      float64       `json:"volume"`
}

func NewCandle(productCode string, duration time.Duration, timeDate time.Time, open, close, high, low, volume float64) *Candle {
	return &Candle{
		productCode,
		duration,
		timeDate,
		open,
		close,
		high,
		low,
		volume,
	}
}

func GetCandle(productCode string, duration time.Duration, dateTime time.Time) *Candle {
	ddl := fmt.Sprintf("SELECT time, open, close, high, low, volume FROM %s WHERE time = ?", GetCandleTableName(productCode, duration))
	row := DbConnection.QueryRow(ddl, dateTime)

	var candle Candle
	err := row.Scan(&candle.Time, &candle.Open, &candle.Close, &candle.High, &candle.Low, &candle.Volume)
	if err != nil {
		return nil
	}

	return NewCandle(productCode, duration, candle.Time, candle.Open, candle.Close, candle.High, candle.Low, candle.Volume)
}

func (c *Candle) GetTableName() string {
	return GetCandleTableName(c.ProductCode, c.Duration)
}

func CreateCandleWithDuration(ticker bitflyer.Ticker, productCode string, duration time.Duration) bool {
	currentCandle := GetCandle(productCode, duration, ticker.TruncateDateTime(duration))
	price := ticker.GetMidPrice()
	if currentCandle == nil {
		candle := NewCandle(productCode, duration, ticker.TruncateDateTime(duration), price, price, price, price, ticker.Volume)
		err := candle.Create()
		if err != nil {
			log.Fatalln(err)
		}
		return true
	}

	if currentCandle.High <= price {
		currentCandle.High = price
	} else if currentCandle.Low >= price {
		currentCandle.Low = price
	}
	currentCandle.Volume += ticker.Volume
	currentCandle.Close = price
	currentCandle.Update()
	return false
}

func (c *Candle) Create() error {
	ddl := fmt.Sprintf("INSERT INTO %s (time, open, close, high, low, volume) VALUES (?, ?, ?, ?, ?, ?)", c.GetTableName())
	_, err := DbConnection.Exec(ddl, c.Time, c.Open, c.Close, c.High, c.Low, c.Volume)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (c *Candle) Update() error {
	ddl := fmt.Sprintf("UPDATE %s SET open = ?, close = ?, high = ?, low = ?, volume = ? WHERE time = ?", c.GetTableName())
	_, err := DbConnection.Exec(ddl, c.Open, c.Close, c.High, c.Low, c.Volume, c.Time)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
