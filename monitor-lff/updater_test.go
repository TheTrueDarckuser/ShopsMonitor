package lff

import (
	"testing"
	"time"

	"/restock-monitor/parser"
)

func TestUpdater(t *testing.T) {
	updates = new(outChannel)

	worker := SpawnUpdateWorker(
		func() ([]*parser.Product, error) {
			return []*parser.Product{&parser.Product{Name: "Bibba"}, &parser.Product{Name: "Bobba"}}, nil
		},
		updates,
		WorkerOptions{
			time.Millisecond * 10,
			make(chan error),
		},
	)

	_ = worker

	// test failes if didn't recieve six products within 5 seconds
	cnt := 0
	timer := time.NewTimer(time.Second * 3)

	for {
		select {
		case product := <-updates:
			t.Logf("received product named %s", product.Name)
			cnt++
			if cnt >= 4 {
				return
			}
		case <-timer.C:
			t.Errorf("time exceeded limit")
			return
		}
	}
}
