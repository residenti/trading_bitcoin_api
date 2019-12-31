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

func init() {
	dbConnection, err := sql.Open(config.List.Driver, config.List.DataSource)
	if err != nil {
		log.Fatalln(err)
	}
	defer dbConnection.Close() // 正直Closeすべきか曖昧.

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
		_, err := dbConnection.Exec(ddl)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func GetCandleTableName(productCode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", strings.ToLower(productCode), duration)
}
