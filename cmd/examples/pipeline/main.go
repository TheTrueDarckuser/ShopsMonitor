package main

import (
	"log"
	"os"
	"time"

	lff "/restock-monitor/monitor-lff"
	"/restock-monitor/parser"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	mgoSess, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	patternsWrapper := lff.NewPatternsWrapper(mgoSess.DB("groc").C("patterns"))
	patternsWrapper.NewPatternsStore()

	productsWrapper := lff.NewProductsWrapper(mgoSess.DB("groc").C("products"))
	productWrappersStore := productsWrapper.NewProductWrappersStore()

	f, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")

	_, cPool, fPool, errors, newlyAvailable := lff.StartMonitoringPipeline(
		4*time.Second,
		parser.Global,
		patternsWrapper,
		productWrappersStore,
	)

	_ = cPool
	_ = fPool

	for {
		select {
		case product := <-newlyAvailable:
			log.Println(product.Name, product.Link)
		case err := <-errors:
			log.Printf("fcked up: %s\n", err.Error())
		}
	}
}
