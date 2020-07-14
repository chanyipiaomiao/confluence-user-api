package main

import (
	"flag"
	"log"
)

func main() {

	cfgPath := flag.String("cfg_path", "./config.json", "config file path")
	logPath := flag.String("log_path", "./confluence-user-api.log", "log path")
	flag.Parse()

	if err := InitLogger(*logPath); err != nil {
		log.Println(err)
		return
	}

	if _, err := InitConfig(*cfgPath); err != nil {
		log.Println(err)
		return
	}

	InitServer()
}
