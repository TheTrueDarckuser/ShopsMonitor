package lff

import (
	"fmt"
	"math/rand"
	"time"

	"/restock-monitor/parser"
)

// CheckerWorkerPool manages CheckerWorkers
type CheckerWorkerPool struct {
	workers         map[string]*CheckerWorker
	website2checker map[string]parser.AvailabilityChecker
	i               inChannel
	o               outChannel
	options         Options
}

// StartCheckerPool starts a new checker pool
func StartCheckerPool(
	website2checker map[string]parser.AvailabilityChecker,
	i inChannel, o outChannel,
	options Options) *CheckerWorkerPool {

	pool := new(CheckerWorkerPool)
	pool.workers = make(map[string]*CheckerWorker)
	pool.website2checker = website2checker
	pool.i = i
	pool.o = o
	pool.options = options

	pool.startProductPulling()

	return pool
}

// startProductPulling pulls product from input channel and monitores them if checker exists
func (wp *CheckerWorkerPool) startProductPulling() {
	go func() {
		for {
			select {
			case pw := <-wp.i:
				if err := wp.MonitorProduct(pw); err != nil {
					wp.options.errors <- err
				}
			}
		}
	}()
}

// MonitorProduct assignes worker from website2checker map to product
func (wp *CheckerWorkerPool) MonitorProduct(pw *ProductWrapper) error {
	if job, ok := wp.website2checker[pw.ShopURL]; ok {
		wp.startWorker(job, pw) // do not start this in goroutine
		return nil
	}

	return fmt.Errorf("checker for %s not found", pw.ShopURL)
}

// StartWorker starts new checker worker
func (wp *CheckerWorkerPool) startWorker(job parser.AvailabilityChecker, pw *ProductWrapper) {
	if _, exists := wp.workers[pw.Link]; exists {
		return
	}
	wp.workers[pw.Link] = SpawnCheckerWorker(
		job,
		pw,
		wp.o,

		Options{
			wp.options.delay + time.Millisecond*time.Duration(rand.Intn(1000)-1000),
			wp.options.errors,
		},
	)
}

// KillWorker kills worker
func (wp *CheckerWorkerPool) KillWorker(link string) error {
	if worker, exists := wp.workers[link]; exists {
		worker.Kill()
		delete(wp.workers, link)
		return nil
	}

	return fmt.Errorf("worker for %s not found", link)
}

// Workers returns all workers in pool
func (wp CheckerWorkerPool) Workers() map[string]*CheckerWorker {
	return wp.workers
}

// ########################################################################
// ########################################################################
// ########################################################################
// ########################################################################

// CheckerWorker performs product specific monitoring
type CheckerWorker struct {
	job parser.AvailabilityChecker // parser checking for product avaiability
	pw  *ProductWrapper            // monitored product

	o      outChannel // on-availability channel (when avaiability state changes from false -> true)
	errors chan error // runtime errors channel
	cancel chan bool  // kills worker

	ticker *time.Ticker // checks performed every %time durations
}

// SpawnCheckerWorker spawnes a new worker and returns its cancellation channel
func SpawnCheckerWorker(job parser.AvailabilityChecker, pw *ProductWrapper, o outChannel, options Options) *CheckerWorker {
	worker := new(CheckerWorker)

	worker.job = job
	worker.pw = pw
	worker.o = o
	worker.errors = options.errors
	worker.cancel = make(chan bool)
	worker.ticker = time.NewTicker(options.delay)

	worker.Start()

	return worker
}

// Start starts monitoring routine
func (cw *CheckerWorker) Start() {
	go func() {
		for {
			select {
			case <-cw.ticker.C:
				state, err := cw.job(cw.pw.Link)
				if err != nil {
					cw.errors <- err
				} else if (cw.pw.State != 1) && (state) {
					cw.pw.State = 1
					cw.o <- cw.pw
				} else if !(state) {
					cw.pw.State = 2
					if cw.pw.State == 0 {
						cw.pw.Type = "coming soon"
						cw.o <- cw.pw
					}
				}

			case <-cw.cancel:
				cw.pw.State = 3
				close(cw.cancel)
				return
			}
		}
	}()
}

// Kill kills worker
func (cw *CheckerWorker) Kill() {
	cw.cancel <- true
}
