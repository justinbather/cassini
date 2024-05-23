package main

import (
	"io"
	"log"
	"net/http"

	c "github.com/justinbather/cassini/pkg/config"
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

	log.Printf("Cassini Config\n ---------\n %s\n ---------\n Port: %d\n URL: %s\n Method: %s\n", config.Service.Name, config.Service.Port, config.Service.Url, config.Service.Method)

	client := http.Client{}

	req, err := http.NewRequest(config.Service.Method, config.Service.Url, nil)
	if err != nil {
		log.Print(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}

	log.Print(string(body))
}
