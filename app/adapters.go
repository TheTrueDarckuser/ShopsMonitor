package main

import (
	"net/http"
	"sync"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	bot "/restock-monitor/bot/discord"
	pm "/restock-monitor/monitor-lff"
	mgo "gopkg.in/mgo.v2"
)

// Adapter is an http adapter
type Adapter func(http.Handler) http.Handler

// Adapt adapts a http handler
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func withProductWrappersStore(productsstore *sync.Map) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, "productWrappersStore", productsstore)
			h.ServeHTTP(w, r)
		})
	}
}

func withFW(fw *pm.KeywordFilter) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, "fw", fw)
			h.ServeHTTP(w, r)
		})
	}
}

func withDB(db *mgo.Database) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, "db", db)
			h.ServeHTTP(w, r)
		})
	}
}

func withAccess(store sessions.CookieStore, onDenie http.HandlerFunc) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if session, err := store.Get(r, "rmonitor"); err == nil {
				if v, ok := session.Values["access"]; ok && v.(bool) == true {
					h.ServeHTTP(w, r)
					return
				}
				onDenie.ServeHTTP(w, r)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	}
}

func withSession(store sessions.CookieStore) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if session, err := store.Get(r, "rmonitor"); err == nil {
				context.Set(r, "session", session)
				h.ServeHTTP(w, r)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
	}
}

func withCW(cw *pm.CheckerWorkerPool) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, "cw", cw)
			h.ServeHTTP(w, r)
		})
	}
}

func withBot(bot *bot.Bot) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, "bot", bot)
			h.ServeHTTP(w, r)
		})
	}
}
