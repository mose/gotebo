package main

import (
	"github.com/tsuru/config"
	"log"
)

var (
	ServerName string
)

func Config() error {
	err := config.ReadConfigFile(ConfigFile)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
