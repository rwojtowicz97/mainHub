package main

import (
	"github.com/rwojtowicz97/mainHub/api"
	"github.com/rwojtowicz97/mainHub/configs"
)

func main() {
	conf := configs.Config{
		Port: ":8080",
	}
	api.Init(conf)
}
