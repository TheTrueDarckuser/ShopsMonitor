package lff

import (
	"sync"
	"testing"

	"/restock-monitor/parser"
)

func TestKeywordFilter(t *testing.T) {

	pwStore := new(sync.Map)
	pStore := new(sync.Map)

	pwStore.Store(
		"google.com", ProductWrapper{&parser.Product{Name: "correct Store"}, 0},
	)
	pwStore.Store(
		"google.com", ProductWrapper{&parser.Product{Name: "correct lil true Store"}, 0},
	)
	pwStore.Store(
		"google.com", ProductWrapper{&parser.Product{Name: "incorrect Store"}, 0},
	)

	pStore.Store("correct", nil)
	pStore.Store("correct&true", nil)

	i := make(chan *ProductWrapper)
	o := make(chan *ProductWrapper)

	filter := NewKeywordFilter(pStore, pwStore, i, o)

	products := []*ProductWrapper{
		{&parser.Product{Name: "correct good chan"}, 0},
		{&parser.Product{Name: "correct true chan"}, 0},
		{&parser.Product{Name: "incorrect incorrect chan"}, 0},
	}

	for _, tt := range products {
		t.Run(tt.Name, func(t *testing.T) {
			var out bool

			i <- tt

			select {
			case <-o:
				out = true
			default:
				out = false
			}

			if tt.out != out {
				t.Errorf("wrong filtering!: %t != %t", tt.out, out)
			}
		})
	}
}
