package lff

import (
	"testing"
	"time"

	"/restock-monitor/parser"
)

func uJob() ([]*parser.Product, error) {
	return []*parser.Product{}, nil
}

func TestUpdateWorkerPool(t *testing.T) {
	e := make(chan error)
	o := make(outChannel)
	pool := StartUpdaterPool(
		map[string]parser.ProductsUpdater{
			"website.com": uJob,
		},
		time.Millisecond*500,
		e,
		o,
	)

	worker := pool.Workers()["website.com"]

	t.Run("adding worker to updater pool", func(t *testing.T) {
		if len(pool.Workers()) != 1 {
			t.Errorf("didn't add worker to pool")
			return
		}
	})

	t.Run("adding existing worker to updater pool", func(t *testing.T) {
		pool.startWorker(
			"website.com",
			uJob,
		)

		if len(pool.Workers()) > 1 {
			t.Errorf("added existing worker to pool")
		}
	})

	t.Run("removing worker", func(t *testing.T) {
		pool.killWorker("website.com")

		if len(pool.Workers()) > 0 {
			t.Errorf("worker didn't get killed")
		}

		if _, ok := <-worker.cancel; ok {
			t.Errorf("worker killed, but still running")
		}
	})
}

func cJob(string) (bool, error) {
	return false, nil
}
func TestCheckerPool(t *testing.T) {
	testChan := make(chan struct{})

	pool := StartCheckerPool(
		map[string]parser.AvailabilityChecker{
			"website.com": func(string) (bool, error) {
				return false, nil
			},

			"test.com": func(string) (bool, error) {
				testChan <- struct{}{}
				return false, nil
			},
		},

		time.Microsecond*30,
		make(chan ProductWrapper),
		make(chan error),
	)

	product := ProductWrapper{
		Product: parser.Product{
			Name:    "test product",
			ShopURL: "website.com",
		},
	}

	t.Run("adding new checker worker", func(t *testing.T) {
		pool.startWorker(cJob, product)

		if len(pool.Workers()) != 1 {
			t.Errorf("didn't add worker to pool")
		}
	})

	t.Run("adding existing worker to checker pool", func(t *testing.T) {
		pool.startWorker(cJob, product)

		if len(pool.Workers()) > 1 {
			t.Errorf("added existing product to pool")
		}
	})

	t.Run("killing worker", func(t *testing.T) {
		worker := pool.Workers()[product]
		pool.killWorker(product)

		if len(pool.Workers()) > 0 {
			t.Errorf("worker not killed")
		}

		if _, ok := <-worker.cancel; ok {
			t.Errorf("worker killed, but still working")
		}
	})

	t.Run("add worker by product", func(t *testing.T) {
		if err := pool.monitorProduct(product); err != nil {
			t.Error(err)
			return
		}

		if len(pool.Workers()) != 1 {
			t.Errorf("didn't or added wrong number of workers to pool, expected 1, got %d", len(pool.Workers()))
		}
	})

	t.Run("test product pulling", func(t *testing.T) {
		pool.I <- ProductWrapper{
			Product: parser.Product{
				Name:    "bibba",
				ShopURL: "test.com",
			},
		}

		select {
		case <-testChan:
			return
		case <-time.NewTimer(time.Second).C:
			t.Errorf("didn't start worker for pulled product, expected 2 workers, got %d", len(pool.Workers()))
		}
	})
}
