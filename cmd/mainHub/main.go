package main

import (
	"net/http"

	"github.com/rwojtowicz97/mainHub/internal"
)

func main() {
	client := http.Client{}

	tw := internal.NewChannel("h2p_gucio", &client)

	tw.CheckIfLive()

	// conf := configs.Config{
	// 	Port: ":3000",
	// }
	// api.Init(conf)
}
