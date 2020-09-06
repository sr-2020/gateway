package main

import (
	"github.com/sr-2020/gateway/app"
	"github.com/sr-2020/gateway/app/adapters/config"
)

func main() {
	cfg := config.LoadConfig()

	_ = app.Start(cfg)
}
