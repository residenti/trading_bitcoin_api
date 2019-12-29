package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Logfile string
}

var List ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Faild to read file: %v", err)
		os.Exit(1)
	}

	List = ConfigList{
		Logfile: cfg.Section("trading_bitcoin_api").Key("log_file").String(),
	}
}
