package lff

import (
	"strings"
	"sync"
)

// KeywordFilter filters products by name
type KeywordFilter struct {
	PStore  *sync.Map
	PWStore *sync.Map
	i       inChannel
	o       outChannel
}

// NewKeywordFilter creates a new KeyWordFilter instance
func NewKeywordFilter(pStore, pwStore *sync.Map, i inChannel, o outChannel) *KeywordFilter {
	filter := new(KeywordFilter)

	filter.PStore = pStore
	filter.PWStore = pwStore
	filter.i = i
	filter.o = o

	filter.PushProductsFromDB()
	filter.Start()

	return filter
}

// PushProductsFromDB pushes products(with state != 3) from db to get them checked again
func (f *KeywordFilter) PushProductsFromDB() {
	go func() {
		f.PWStore.Range(func(k, v interface{}) bool {
			if pw := v.(*ProductWrapper); pw.State != 3 {
				f.o <- pw
			}
			return true
		})
	}()
}

// Start starts keyword filter
func (f *KeywordFilter) Start() {
	go func() {
		for {
			select {
			case pw := <-f.i:
				if f.isMatch(pw.Name) {
					if _, loaded := f.PWStore.LoadOrStore(pw.Link, pw); !loaded {
						f.o <- pw
					}
				}
			}
		}
	}()
}

// isMatch tells if a product is urgent
func (f *KeywordFilter) isMatch(name string) bool {
	ok := false
	f.PStore.Range(func(k, v interface{}) bool {
		if contains(name, k.(string)) {
			ok = true
		}
		return true
	})
	return ok
}

// AddPattern adds new pattern to filter
func (f *KeywordFilter) AddPattern(pattern string) {
	for _, p := range strings.Split(pattern, "&&") {
		f.PStore.LoadOrStore(strings.ToLower(strings.Replace(p, " ", "", -1)), nil)
	}
}

// RemovePattern removes pattern from filter
func (f *KeywordFilter) RemovePattern(pattern string) {
	f.PStore.Delete(pattern)
}

func contains(name, pattern string) bool {
	keys := strings.Split(pattern, "&")
	for _, key := range keys {
		if !strings.Contains(name, key) {
			return false
		}
	}
	return true

}
