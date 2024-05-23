package main

import (
	"fmt"
	"log"

	"github.com/justinbather/cassini/pkg/types"

	"github.com/spf13/viper"
)

func main() {

	fmt.Println("Cassini")

	viper.SetConfigFile("./test.yaml")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config types.CassiniConfig

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	log.Println(config)

	fmt.Printf("Cassini Config\n Port: %d\n URL: %s\n Method: %s\n", config.Server.Port, config.Server.Url, config.Server.Method)
}
