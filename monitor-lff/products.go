package lff

import (
	"sync"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// ProductsWrapper contains stuff for synchronizing products and db
type ProductsWrapper struct {
	C                    *mgo.Collection
	ProductWrappersStore *sync.Map
}

// NewProductsWrapper return new productsWrapper
func NewProductsWrapper(C *mgo.Collection) *ProductsWrapper {
	productsWrapper := new(ProductsWrapper)
	productsWrapper.C = C
	productsWrapper.ProductWrappersStore = new(sync.Map)
	return productsWrapper
}

// NewProductWrappersStore returns new products store
func (pw *ProductsWrapper) NewProductWrappersStore() *sync.Map {
	var productWrappers []*ProductWrapper
	pw.C.Find(nil).All(&productWrappers)
	for _, productWrapper := range productWrappers {
		pw.ProductWrappersStore.Store(productWrapper.Link, productWrapper)
	}
	return pw.ProductWrappersStore
}

// StartSync starts the automatic synchronization with db
func (pw *ProductsWrapper) StartSync() {
	for {
		time.Sleep(8 * time.Hour)
		pw.ApplyStoreToDB()
	}
}

// ApplyStoreToDB applies all changes to db
func (pw *ProductsWrapper) ApplyStoreToDB() {
	pw.C.DropCollection()
	bulk := pw.C.Bulk()
	pw.ProductWrappersStore.Range(func(k, v interface{}) bool {
		bulk.Insert(v)
		return true
	})
	bulk.Run()
}
