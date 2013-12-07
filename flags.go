package main

import (
	"flag"
)

var (
	ConfigFile string
)

func init() {
	flag.StringVar(&ConfigFile, "config", "gotebo.yml", "Path to the getebo config file")
}

func Flags() error {
	if !flag.Parsed() {
		flag.Parse()
	}
	return nil
}
