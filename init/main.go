package main

import (
	"flag"
	"web-study/init/cmd"
)

var configFileFlag = flag.String("config", "config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configFileFlag)
}
