package main

import (
	"context"

	"github.com/peanut-pg/gin_admin/internal/app"
	"github.com/peanut-pg/gin_admin/pkg/logger"
)

var VERSION = "0.0.1"

func main() {
	logger.SetVersion(VERSION)
	ctx := logger.NewTraceIDContext(context.Background(), "main")
	app.Run(ctx,
		app.SetVersion(VERSION),
		app.SetConfigFile("configs/config.toml"))

}
