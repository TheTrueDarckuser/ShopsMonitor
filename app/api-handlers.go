package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/context"
	bot "/restock-monitor/bot/discord"
	pm "/restock-monitor/monitor-lff"
	pp "/restock-monitor/parser"
)

type shopdata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func handleGetShopURLS(w http.ResponseWriter, r *http.Request) {
	var data []shopdata
	for k, v := range pp.Global {
		data = append(data, shopdata{k, v.ShopName})
	}
	bytes, _ := json.Marshal(data)
	w.Write(bytes)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetProducts(w, r)
	case "DELETE":
		handleIgnoreProduct(w, r)
	}
}

func handleGetProducts(w http.ResponseWriter, r *http.Request) {
	productWrappersStore := context.Get(r, "productWrappersStore").(*sync.Map)

	var products []pp.Product
	productWrappersStore.Range(func(k, v interface{}) bool {
		if k := v.(*pm.ProductWrapper).State; k == 1 || k == 2 {
			products = append(products, *v.(*pm.ProductWrapper).Product)
		}
		return true
	})

	if bytes, err := json.Marshal(products); err == nil {
		fmt.Fprint(w, string(bytes))
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleIgnoreProduct(w http.ResponseWriter, r *http.Request) {
	var product pp.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fw := context.Get(r, "fw").(*pm.KeywordFilter)
	cw := context.Get(r, "cw").(*pm.CheckerWorkerPool)
	if _, ok := fw.PWStore.Load(product.Link); ok {
		cw.KillWorker(product.Link)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func patternsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetPatterns(w, r)
	case "POST":
		handleAddPattern(w, r)
	case "DELETE":
		handleDeletePattern(w, r)
	}
}

func handleGetPatterns(w http.ResponseWriter, r *http.Request) {
	fw := context.Get(r, "fw").(*pm.KeywordFilter)

	var patterns []string
	fw.PStore.Range(func(k, v interface{}) bool {
		patterns = append(patterns, k.(string))
		return true
	})

	if bytes, err := json.Marshal(patterns); err == nil {
		fmt.Fprint(w, string(bytes))
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleAddPattern(w http.ResponseWriter, r *http.Request) {
	var pattern pm.Pattern
	err := json.NewDecoder(r.Body).Decode(&pattern)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fw := context.Get(r, "fw").(*pm.KeywordFilter)
	fw.AddPattern(pattern.Pattern)
}

func handleDeletePattern(w http.ResponseWriter, r *http.Request) {
	var pattern pm.Pattern
	err := json.NewDecoder(r.Body).Decode(&pattern)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fw := context.Get(r, "fw").(*pm.KeywordFilter)
	fw.RemovePattern(pattern.Pattern)
}

func handleAddProduct(w http.ResponseWriter, r *http.Request) {
	var product pp.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err == nil {
		cw := context.Get(r, "cw").(*pm.CheckerWorkerPool)
		fw := context.Get(r, "fw").(*pm.KeywordFilter)
		pw := &pm.ProductWrapper{Product: &product, State: 0}
		if v, loaded := fw.PWStore.LoadOrStore(product.Link, pw); !loaded || v.(*pm.ProductWrapper).State == 2 {
			cw.MonitorProduct(pw)
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func discordSettingsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetDiscord(w, r)
	case "POST":
		handlePostDiscord(w, r)
	case "PATCH":
		handlePatchDiscord(w, r)
	case "DELETE":
		handleDeleteDiscord(w, r)
	}
}

func handleGetDiscord(w http.ResponseWriter, r *http.Request) {
	bot := context.Get(r, "bot").(*bot.Bot)

	channels, err := bot.Session.GuildChannels(bot.Guild.ID)
	if err != nil {
		return
	}

	if bytes, err := json.Marshal(channels); err == nil {
		fmt.Fprint(w, string(bytes))
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlePatchDiscord(w http.ResponseWriter, r *http.Request) {

}

func handlePostDiscord(w http.ResponseWriter, r *http.Request) {

}

func handleDeleteDiscord(w http.ResponseWriter, r *http.Request) {

}
