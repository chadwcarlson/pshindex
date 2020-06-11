package main

import (
	"github.com/chadwcarlson/gomeili/index"
	"github.com/chadwcarlson/gomeili/utils/config"
)

func main() {

	var config config.ConfigSet
	combinedIndex := index.Build(config.Load("config.yaml"))
	combinedIndex.Write("output/pshindex.json")

}
