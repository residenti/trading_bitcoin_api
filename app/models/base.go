package models

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/residenti/trading_bitcoin_api/config"
)

var DbConnection *sql.DB

func init() {
	var err error
	DbConnection, err = sql.Open(config.List.Driver, config.List.DataSource)
	if err != nil {
		log.Fatalln(err)
	}

	// NOTE Don't do that if you don't want your db closed when NewDatabase returns. You don't need to close the db if you plan on reusing it. However you need to close rows whenever you call Query, otherwise your app will hit the connection limit and crash.
	// defer DbConnection.Close() // 正直Closeすべきか曖昧.

	for _, duration := range config.List.Durations {
		tableName := GetCandleTableName(config.List.ProductCode, duration)
		ddl := fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				time DATETIME PRIMARY KEY NOT NULL,
				open FLOAT,
				close FLOAT,
				high FLOAT,
				low FLOAT,
				volume FLOAT)`, tableName)
		_, err := DbConnection.Exec(ddl)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func GetCandleTableName(productCode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", strings.ToLower(productCode), duration)
}
