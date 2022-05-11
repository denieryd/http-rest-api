package main

import (
    "flag"
    "github.com/BurntSushi/toml"
    "github.com/denieryd/http-rest-api/internal/app/apiserver"
    "log"
)

var configPath string

func init() {
    flag.StringVar(&configPath, "config", "configs/apiserver.toml", "path to config file")
}

func main() {
    flag.Parse()
    config := apiserver.NewConfig()

    _, err := toml.DecodeFile(configPath, config)
    if err != nil {
        log.Fatal(err)
    }

    server := apiserver.New(config)
    if err := server.Start(); err != nil {
        log.Fatal(err)
    }

}
