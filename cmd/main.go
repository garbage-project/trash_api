package main

import (
	"flag"
	"fmt"

	"github.com/garbage-project/trash_api.git/cmd/app"
	"github.com/garbage-project/trash_api.git/config"
)

var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

const (
	banner = `

`
)

func main() {
	flag.Parse()
	cfg := config.NewConfig(*configFlag)
	fmt.Printf(banner)
	a := app.NewApp(cfg)
	go a.Wait()
	a.Run()
}
