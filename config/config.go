package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Item string
}

var List ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Faild to read file: %v", err)
		os.Exit(1)
	}

	List = ConfigList{
		Item: cfg.Section("config").Key("item").String(),
	}
}
