package lff

import (
	"sync"
	"time"

	"/restock-monitor/parser"
)

type inChannel <-chan *ProductWrapper
type outChannel chan<- *ProductWrapper

// StartMonitoringPipeline starts full monitoring pipeline
func StartMonitoringPipeline(
	delay time.Duration,
	wuc map[string]parser.GlobalWrapper,
	pStore, pwStore *sync.Map) (*UpdateWorkerPool, *CheckerWorkerPool, *KeywordFilter, chan error, chan *ProductWrapper) {

	errors := make(chan error)
	newArrivals := make(chan *ProductWrapper)
	filtered := make(chan *ProductWrapper)
	newlyAvailable := make(chan *ProductWrapper)
	options := Options{delay: delay, errors: errors}

	uPool := StartUpdaterPool(
		getW2U(),
		newArrivals,
		options,
	)

	filter := NewKeywordFilter(
		pStore,
		pwStore,
		newArrivals,
		filtered,
	)

	cPool := StartCheckerPool(
		getW2C(),
		filtered,
		newlyAvailable,
		options,
	)

	return uPool, cPool, filter, errors, newlyAvailable
}
