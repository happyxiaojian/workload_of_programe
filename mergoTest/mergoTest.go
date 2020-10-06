package main

import (
	"fmt"
	"log"

	"github.com/imdario/mergo"
)

type RedisConfig struct {
	Address string
	Port int
	DB int
}

var defaultConfig = RedisConfig{
	Address:"127.0.0.1",
	Port: 6381,
	DB:1,
}

func main(){
	var config RedisConfig
	var m = make(map[string]interface{})

	if err := mergo.Map(&m, defaultConfig); err != nil {
		log.Fatal(err)
	}


	config = defaultConfig

	defaultConfig.DB = 2

	fmt.Printf("config:%#v\n", config)
	fmt.Printf("config:%#v\n", m)
}
