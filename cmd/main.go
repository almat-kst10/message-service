package main

import (
	"log"

	"github.com/almat-kst10/message-service/configs"
	"github.com/almat-kst10/message-service/internal/app"
)

func main() {
	config, err := configs.NewConfigs()
	if err != nil {
		log.Println("err", err)
		return
	}

	err = app.Run(config)
	if err != nil {
		return
	}
}
