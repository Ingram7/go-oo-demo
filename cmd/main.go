package main

import (
	"context"
	"go-oo-demo/internal/app"
	"go-oo-demo/pkg/logger"
)

// VERSION 版本号，可以通过编译的方式指定版本号：go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "1.0.0"

func main() {

	logger.SetVersion(VERSION)
	ctx := logger.NewTagContext(context.Background(), "__main__")

	engine := new(app.Engine)
	engine.Start(ctx)
}