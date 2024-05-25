package main

import (
	"log"

	c "github.com/justinbather/cassini/pkg/config"
	"github.com/justinbather/cassini/pkg/runner"
	"github.com/justinbather/cassini/pkg/ticker"
	"github.com/spf13/viper"
)

func main() {

	log.Println("Cassini")

	viper.SetConfigFile("./test.yaml")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config c.CassiniConfig

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	log.Printf("Cassini Config\n ---------\n %s\n ---------\n Port: %d\n URL: %s\n", config.Service.Name, config.Service.Port, config.Service.Url)
	log.Println("Tests:", config.Service.Tests)

	ticker := ticker.BuildTicker(config)

	exit := make(chan bool)

	for {

		go runner.Run(exit, ticker, config)
	}

}
