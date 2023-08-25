package lff

import (
	"sync"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// PatternsWrapper contains stuff for synchronizing patterns and patterns
type PatternsWrapper struct {
	C             *mgo.Collection
	PatternsStore *sync.Map
}

// Pattern represents pattern
type Pattern struct {
	Pattern string `json:"pattern" bson:"pattern"`
}

// NewPatternsWrapper return new patternsWrapper
func NewPatternsWrapper(C *mgo.Collection) *PatternsWrapper {
	patternsWrapper := new(PatternsWrapper)
	patternsWrapper.C = C
	patternsWrapper.PatternsStore = new(sync.Map)
	return patternsWrapper
}

// NewPatternsStore returns new patterns store
func (pw *PatternsWrapper) NewPatternsStore() *sync.Map {
	var patterns []*Pattern
	pw.C.Find(nil).All(&patterns)
	for _, pattern := range patterns {
		pw.PatternsStore.Store(pattern.Pattern, nil)
	}
	return pw.PatternsStore
}

// StartSync starts the automatic synchronization with db
func (pw *PatternsWrapper) StartSync() {
	for {
		time.Sleep(8 * time.Hour)
		pw.ApplyStoreToDB()
	}
}

// ApplyStoreToDB applies all changes to db
func (pw *PatternsWrapper) ApplyStoreToDB() {
	pw.C.DropCollection()
	bulk := pw.C.Bulk()
	pw.PatternsStore.Range(func(k, v interface{}) bool {
		bulk.Insert(&Pattern{Pattern: k.(string)})
		return true
	})
	bulk.Run()
}
