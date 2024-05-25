package ticker

import (
	"log"
	"time"

	"github.com/justinbather/cassini/pkg/config"
)

func BuildTicker(config config.CassiniConfig) time.Ticker {
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

	return ticker
}
