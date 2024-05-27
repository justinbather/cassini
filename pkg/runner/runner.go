package runner

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/justinbather/cassini/pkg/config"
)

func Run(exit chan bool, ticker time.Ticker, config config.CassiniConfig) {
	client := http.Client{}

	tests := config.Service.Tests

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

				if resp.StatusCode != test.AssertStatus {
					log.Printf("\n%s failed. expected response status %d but got %d", test.Name, test.AssertStatus, resp.StatusCode)
				} else {
					log.Printf("\nPASS: %s\nStatus Expected: %d\nStatus Received: %d", test.Name, test.AssertStatus, resp.StatusCode)
				}

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Print(err)
				}

				log.Print(string(body))
			}
		}
	}

}
