package models

import (
	"fmt"
	"time"
)

type DataFrameCandle struct {
	ProductCode string        `json:"product_code"`
	Duration    time.Duration `json:"duration"`
	Candles     []Candle      `json:"candles"`
}

func GetDataFrameCandle(productCode string, duration time.Duration, limit int) (dfCandle *DataFrameCandle, err error) {
	// NOTE AS でエイリアス名を付けないとエラーになる. Error 1248: Every derived table must have its own alias
	ddl := fmt.Sprintf(`SELECT * FROM (
		SELECT time, open, close, high, low, volume FROM %s ORDER BY time DESC limit ?
	) AS btc_jpy ORDER BY time ASC;`, GetCandleTableName(productCode, duration))
	rows, err := DbConnection.Query(ddl, limit)
	if err != nil {
		return
	}
	defer rows.Close()

	dfCandle = &DataFrameCandle{}
	dfCandle.ProductCode = productCode
	dfCandle.Duration = duration

	for rows.Next() {
		var candle Candle

		candle.ProductCode = productCode
		candle.Duration = duration
		rows.Scan(&candle.Time, &candle.Open, &candle.Close, &candle.High, &candle.Low, &candle.Volume)

		dfCandle.Candles = append(dfCandle.Candles, candle)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return dfCandle, nil
}
