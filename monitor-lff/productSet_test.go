package lff

import (
	"testing"

	"/restock-monitor/parser"
)

func TestProductSet(t *testing.T) {
	pset := newProductSet()

	products := []parser.Product{
		{Name: "sneaker", Link: "shop.com/abc"},
		{Name: "sneaker 2", Link: "shop.com/abc"},
		{Name: "sneaker", Link: "shop.cim/def"},
	}

	for _, p := range products {
		pset.Add(ProductWrapper{
			Product: p,
		})
	}

	t.Run("getting existing product", func(t *testing.T) {
		if _, found := pset.Get("shop.com/abc"); !found {
			t.Errorf("didn't find added product")
		}
	})

	t.Run("getting non existing product", func(t *testing.T) {
		if _, found := pset.Get("shop.com/ghi"); found {
			t.Errorf("found existing product")
		}
	})

	t.Run("unloading product set", func(t *testing.T) {
		if len(pset.Unload()) != 2 {
			t.Errorf("received unsatisfying nubmer of products, expected 2, received %d", len(pset.Slice()))
		}
	})

	t.Run("checking unloading 2", func(t *testing.T) {
		if len(pset.Slice()) > 0 {
			t.Errorf("didn't dump products afrer unloading")
		}
	})
}
