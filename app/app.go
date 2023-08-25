package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	bot "/restock-monitor/bot/discord"
	pm "/restock-monitor/monitor-lff"
	"gopkg.in/mgo.v2"
)

var (
	headersOk = handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk = handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk = handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	cookieStore *sessions.CookieStore
	shortener   *urlShortener
)

func init() {
	cookieStore = sessions.NewCookieStore([]byte("EF495401D79526606BC45A351DED18E661333B2013451B06153C6CB8ACB88962"))
	cookieStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
	}
}

// c - map,

// StartServer starts the web server
func StartServer(cw *pm.CheckerWorkerPool, fw *pm.KeywordFilter, productWrappersStore *sync.Map, db *mgo.Database, bot *bot.Bot, port string) {
	router := mux.NewRouter()

	// ################
	// # API HANDLERS #
	// ################
	router.Handle(
		"/api/products/",
		Adapt(http.HandlerFunc(productsHandler), withAccess(*cookieStore, accessDenied), withProductWrappersStore(productWrappersStore), withCW(cw), withFW(fw)),
	).Methods("GET", "DELETE")

	router.Handle(
		"/api/auth/",
		Adapt(http.HandlerFunc(handleAuth), withSession(*cookieStore), withDB(db)),
	).Methods("POST")

	router.Handle(
		"/api/join/",
		Adapt(http.HandlerFunc(handleJoin), withSession(*cookieStore), withDB(db)),
	).Methods("POST")

	router.Handle(
		"/api/getShopURLS/",
		Adapt(http.HandlerFunc(handleGetShopURLS), withAccess(*cookieStore, accessDenied)),
	).Methods("GET")

	router.Handle(
		"/api/discord/",
		Adapt(http.HandlerFunc(discordSettingsHandler), withAccess(*cookieStore, accessDenied), withBot(bot)),
	).Methods("GET", "POST", "DELETE")

	router.Handle(
		"/api/patterns/",
		Adapt(http.HandlerFunc(patternsHandler), withAccess(*cookieStore, accessDenied), withFW(fw)),
	).Methods("GET", "POST", "PATCH", "DELETE")

	router.Handle(
		"/api/product/add/",
		Adapt(http.HandlerFunc(handleAddProduct), withAccess(*cookieStore, accessDenied), withCW(cw), withFW(fw)),
	).Methods("POST")

	// #####################
	// # FRONTEND HANDLERS #
	// #####################
	router.Handle(
		"/products",
		Adapt(http.HandlerFunc(productsFrontHandler), withAccess(*cookieStore, authRedirect)),
	).Methods("GET")

	router.Handle(
		"/settings",
		Adapt(http.HandlerFunc(settingsHandler), withAccess(*cookieStore, authRedirect)),
	).Methods("GET")

	router.Handle(
		"/discord",
		Adapt(http.HandlerFunc(discordHandler), withAccess(*cookieStore, authRedirect)),
	).Methods("GET")

	router.Handle(
		"/addProduct",
		Adapt(http.HandlerFunc(addProductHandler), withAccess(*cookieStore, authRedirect)),
	).Methods("GET")

	router.Handle(
		"/auth",
		Adapt(http.HandlerFunc(authHandler), withSession(*cookieStore)),
	).Methods("GET")

	router.Handle(
		"/join",
		Adapt(http.HandlerFunc(joinHandler), withSession(*cookieStore)),
	).Methods("GET")

	router.Handle(
		"/logout",
		Adapt(http.HandlerFunc(logoutHandler), withSession(*cookieStore)),
	).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(authRedirect)

	// static handler
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	log.Printf("Site is running on %s\n", port)
	go func() { http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(router)) }()
}
