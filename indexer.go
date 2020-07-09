package main

import (
	"github.com/chadwcarlson/gomeili/config"
	"github.com/chadwcarlson/gomeili/index"
	"github.com/chadwcarlson/gomeili/server"
	configmeili "github.com/chadwcarlson/gomeili/server/config"
)

func main() {

	var indexConfig config.ConfigSet
	fullIndex := index.Build(indexConfig.Load("config/index.yaml"), "output/combined.json")

	var serverConfig configmeili.Config
	server.Post(serverConfig.Load("config/meili.yaml"), fullIndex)

}
