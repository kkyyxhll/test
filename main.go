package main

import (
	"exchangeapp/config"
	"exchangeapp/router"
)

type Mode struct {
	Iddddddd uint `json:"Id"`
	SEX      string
	NAME     string
}

func main() {
	config.InitConfig()
	r := router.SetupRouter()
	port := config.AppConfig.App.Port
	if port == "" {
		port = "8080"
	}
	r.Run(port)
}
