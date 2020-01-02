package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Logfile       string
	ProductCode   string
	Port          int
	HttpBaseUrl   string
	Durations     map[string]time.Duration
	TradeDuration time.Duration
	Driver        string
	DataSource    string
}

var List ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Faild to read file: %v", err)
		os.Exit(1)
	}

	durations := map[string]time.Duration{
		"1s": time.Second,
		"1m": time.Minute,
		"1h": time.Hour,
	}

	List = ConfigList{
		Logfile:       cfg.Section("trading_bitcoin_api").Key("log_file").String(),
		ProductCode:   cfg.Section("trading_bitcoin_api").Key("product_code").String(),
		Port:          cfg.Section("trading_bitcoin_api").Key("port").MustInt(),
		HttpBaseUrl:   cfg.Section("bitflyer").Key("http_base_url").String(),
		Durations:     durations,
		TradeDuration: durations[cfg.Section("trading_bitcoin_api").Key("trade_duration").String()],
		Driver:        cfg.Section("db").Key("driver").String(),
		DataSource:    cfg.Section("db").Key("data_source").String(),
	}
}
