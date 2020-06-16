package main

import (
	"github.com/chadwcarlson/gomeili/config"
	"github.com/chadwcarlson/gomeili/index"
)

func main() {

	var config config.ConfigSet
	index.Build(config.Load("config.yaml"), "output/combined.json")

}
