package lff

import (
	"sync"
	"testing"
	"time"

	"/restock-monitor/parser"
)

func TestPipeline(t *testing.T) {
	wuc := map[string]parser.GlobalWrapper{
		"website.com": parser.GlobalWrapper{
			Updater: func() ([]*parser.Product, error) {
				return []*parser.Product{&parser.Product{Name: "cool sneaker", ShopURL: "website.com"}}, nil
			},
			Checker: func(string) (bool, error) {
				return true, nil
			},
			ShopName: "website.com",
			Basket:   "basket",
		},
	}

	m := new(sync.Map)

	mm := new(PatternsWrapper)
	mm.PatternsStore.Load("cool&sneaker")

	_, cPool, _, _, _ := StartMonitoringPipeline(
		time.Microsecond*30,
		wuc,
		mm,
		m,
	)

	select {
	case product := <-cPool.i:
		t.Logf("successfully received product %s", product.Name)
	case <-time.NewTimer(time.Second * 3).C:
		t.Errorf("time exceeded")
	}
}
