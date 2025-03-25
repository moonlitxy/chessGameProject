package main

import (
	"common/config"
	"flag"
	"fmt"
)

var configFile = flag.String("config", "application.yml", "config file")

func main() {
	flag.Parse()
	config.InitConfig(*configFile)
	fmt.Println("配置文件信息:", config.Conf)
}
