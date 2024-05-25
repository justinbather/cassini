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

	log.Printf("Cassini Config\n ---------\n %s\n ---------\n Port: %d\n URL: %s\n", config.Service.Name, config.Service.Port, config.Service.Url)
	log.Println("Tests:", config.Service.Tests)

	client := http.Client{}

	tests := config.Service.Tests

	// TODO: We need to loop through the tests and make requests for each one, lets not stop if one fails
	// TODO: We can also change the yaml file to hold the method in each test aswell
	for _, test := range tests {
		req, err := http.NewRequest(test.Method, config.Service.Url, nil)
		if err != nil {
			log.Print(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Print(err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != test.Status {
			log.Printf("\n%s failed. expected response status %d but got %d", test.Name, test.Status, resp.StatusCode)
		} else {
			log.Printf("\nPASS: %s\nStatus Expected: %d\nStatus Received: %d", test.Name, test.Status, resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Print(err)
		}

		log.Print(string(body))
	}
}
