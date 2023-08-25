package lff

import (
	"fmt"
	"time"

	"/restock-monitor/parser"
)

// UpdateWorkerPool manages workers
type UpdateWorkerPool struct {
	workers map[string]*UpdateWorker
	o       outChannel
	options Options
}

// StartUpdaterPool starts a new updater pool
func StartUpdaterPool(jobs map[string]parser.ProductsUpdater, o outChannel, options Options) *UpdateWorkerPool {
	pool := new(UpdateWorkerPool)

	pool.o = o
	pool.workers = make(map[string]*UpdateWorker)
	pool.options = options

	go func() {
		for website, job := range jobs {
			pool.startWorker(website, job)
		}
	}()

	return pool
}

func (wp *UpdateWorkerPool) startWorker(website string, job parser.ProductsUpdater) {
	if _, exists := wp.workers[website]; exists {
		return
	}
	wp.workers[website] = SpawnUpdateWorker(
		job,
		wp.o,
		wp.options,
	)
}

func (wp *UpdateWorkerPool) killWorker(website string) error {
	if worker, exists := wp.workers[website]; exists {
		worker.Kill()
		delete(wp.workers, website)
		return nil
	}
	return fmt.Errorf("worker for %s not found", website)
}

// ########################################################################
// ########################################################################
// ########################################################################
// ########################################################################

// UpdateWorker runs product updates job
type UpdateWorker struct {
	job    parser.ProductsUpdater
	ticker *time.Ticker
	o      outChannel
	errors chan error
	cancel chan bool
}

// SpawnUpdateWorker starts update worker
func SpawnUpdateWorker(job parser.ProductsUpdater, o outChannel, options Options) *UpdateWorker {
	uw := new(UpdateWorker)

	uw.job = job
	uw.o = o
	uw.errors = options.errors
	uw.ticker = time.NewTicker(options.delay)
	uw.cancel = make(chan bool)

	uw.Start()

	return uw
}

// Start starts update worker goroutine
func (uw UpdateWorker) Start() {
	go func() {
		for {
			select {
			case <-uw.ticker.C:
				products, err := uw.job()
				if err != nil {
					uw.errors <- err
				}

				for _, product := range products {
					uw.o <- &ProductWrapper{Product: product, State: 0}
				}
			case <-uw.cancel:
				close(uw.cancel)
				return
			}
		}
	}()
}

// Kill kills updates worker
func (uw *UpdateWorker) Kill() {
	uw.cancel <- true
}
