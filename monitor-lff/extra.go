package lff

import (
	"time"

	"/restock-monitor/parser"
)

// ProductWrapper is a Product wrapper for checker
type ProductWrapper struct {
	*parser.Product `json:"product" bson:"product"`
	State           int `json:"state" bson:"state"` // 0 - unset(just uploaded), 1 - available, 2 - unavailable, 3 - ignored
}

// Options is a struct contaning options
type Options struct {
	delay  time.Duration
	errors chan error
}

// merge merges two product wrapper channels channels
func merge(channels ...chan ProductWrapper) chan ProductWrapper {
	out := make(chan ProductWrapper)

	for _, c := range channels {
		go func() {
			for elm := range c {
				out <- elm
			}
		}()
	}

	return out
}

func getW2U() map[string]parser.ProductsUpdater {
	w2u := make(map[string]parser.ProductsUpdater)
	for k, v := range parser.Global {
		w2u[k] = v.Updater
	}
	return w2u
}

func getW2C() map[string]parser.AvailabilityChecker {
	w2c := make(map[string]parser.AvailabilityChecker)
	for k, v := range parser.Global {
		w2c[k] = v.Checker
	}
	return w2c
}
