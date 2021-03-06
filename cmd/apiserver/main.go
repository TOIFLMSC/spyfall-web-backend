package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/TOIFLMSC/spyfall-web-backend/internal/app/apiserver"
	"github.com/joho/godotenv"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	e := godotenv.Load()
	if e != nil {
		log.Fatal(e)
	}

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
