package main

import (
	"io"
	"log"
	"net/http"
	"time"

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

	var ticker time.Ticker
	switch config.Service.IntervalUnit {
	case "minute":
		ticker = *time.NewTicker(time.Duration(config.Service.IntervalAmount) * time.Minute)
	case "hour":
		ticker = *time.NewTicker(time.Duration(config.Service.IntervalAmount) * time.Hour)
	case "second":
		ticker = *time.NewTicker(time.Duration(config.Service.IntervalAmount) * time.Second)
	default:
		log.Fatal("Must provide a valid interval unit in config.\nThe options are:\n> hour\n> minute\n> second")
	}

	exit := make(chan bool)

	go func() {

		for {
			select {
			case <-exit:
				return
			case t := <-ticker.C:
				log.Println("Tick at ", t)

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
		}

	}()

	time.Sleep(5 * time.Minute)
	ticker.Stop()
	exit <- true
	log.Println("Stopping")
}
