package api

import (
	"github.com/rwojtowicz97/mainHub/configs"
)

func Init(config configs.Config) {
	r := NewRouter()
	r.Run(config.Port)
}
