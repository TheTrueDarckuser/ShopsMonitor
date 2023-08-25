package lff

import (
	"testing"
	"time"

	"/restock-monitor/parser"
)

func TestChecker(t *testing.T) {
	onAvailable := make(chan bool)

	go func() {
		time.Sleep(time.Millisecond)
		onAvailable <- true
	}()

	productChan := make(chan *ProductWrapper)

	worker := SpawnCheckerWorker(
		func(string) (bool, error) {
			select {
			case <-onAvailable:
				return true, nil
			default:
				return false, nil
			}
		},
		&ProductWrapper{
			Product: &parser.Product{Name: "test product"},
			State:   0,
		},
		productChan,
		Options{
			time.Millisecond * 30,
			make(chan error),
		},
	)

	_ = worker

	select {
	case <-time.NewTimer(time.Second * 5).C:
		t.Errorf("time expired")
	case <-productChan:
		t.Log("on time")
	}

	onAvailable <- true
	select {
	case <-time.NewTimer(time.Second).C:
		t.Log("all good")
	case <-productChan:
		t.Error("received already received produt")
	}
}
