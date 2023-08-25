package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	bot "/restock-monitor/bot/discord"
	pm "/restock-monitor/monitor-lff"
	pp "/restock-monitor/parser"
	mgo "gopkg.in/mgo.v2"
)

var (
	info, _      = mgo.ParseURL("mongodb://Ogonyok:Ogonyok1337@ds046037.mlab.com:46037/rmonitor")
	session, err = mgo.DialWithInfo(info)
	db           = session.DB("rmonitor")

	usedKeys = []string{
		"https://www.mrporter.com", "https://www.net-a-porter.com",
		"https://www.ssense.com", "https://www.ssense.com/en-ca/", "https://www.ssense.com/en-gb/",
		"https://www.footlocker.co.uk", "https://www.citygear.com", "https://www.dtlr.com", "https://www.finishline.com",
		"https://www.bstn.com",
		"https://www.footpatrol.com",
		"https://www.sneakersnstuff.com",
	}
)

var (
	port            *string
	discordBotToken *string
	guildID         *string
)

func main() {
	if err != nil {
		log.Fatalf("Couldn't connect to mlab: %s", err)
	}

	// #########
	// # FlAGS #
	// #########
	port = flag.String("port", ":8000", "port server will be running on")
	discordBotToken = flag.String("discordBotToken", "NDk2MDcxMjI5OTE3MTAyMTAx.Dr7fQQ.a9oSBU8XMYr3S6qT5shDz7L68p4", "discord bot token")
	guildID = flag.String("guildID", "413763893043789824", "main guild id")
	flag.Parse()

	// #####################
	// # PATTERNS&PRODUCTS #
	// #####################
	pWrapper := pm.NewPatternsWrapper(db.C("patterns"))
	pStore := pWrapper.NewPatternsStore()

	productsWrapper := pm.NewProductsWrapper(db.C("products"))
	pwStore := productsWrapper.NewProductWrappersStore()

	_, cPool, fPool, errors, newlyAvailable := pm.StartMonitoringPipeline(
		4*time.Second,
		pp.Global,
		pStore,
		pwStore,
	)

	// ###############
	// # DISCORD BOT #
	// ###############
	dbot := bot.NewBot(*discordBotToken, *guildID, newlyAvailable)
	createDiscordChannels(dbot)
	dbot.ListenProductsChannel()

	// #########
	// # FRONT #
	// #########
	go StartServer(cPool, fPool, pwStore, db, dbot, *port)

	// #########
	// # EXTRA #
	// #########
	go pWrapper.StartSync()
	go productsWrapper.StartSync()

	go func() {
		select {
		case e := <-errors:
			log.Println(e.Error())
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	select {
	case <-c:
		ex(dbot, pWrapper, productsWrapper)
		os.Exit(1)
	}
}

func createDiscordChannels(dbot *bot.Bot) {
	var keys []string
	for k := range pp.Global {
		if checkKey(k) {
			keys = append(keys, k)
		}
	}

	//dbot.CreateChannel("grocstuff", keys...)
	dbot.CreateChannel("aio", keys...)
	dbot.CreateChannel("mrp-nap", "https://www.mrporter.com", "https://www.net-a-porter.com")
	dbot.CreateChannel("ssense", "https://www.ssense.com", "https://www.ssense.com/en-ca/", "https://www.ssense.com/en-gb/")
	dbot.CreateChannel("footsites", "https://www.footlocker.co.uk", "https://www.citygear.com", "https://www.dtlr.com", "https://www.finishline.com")
	dbot.CreateChannel("bstn", "https://www.bstn.com")
	dbot.CreateChannel("mesh", "https://www.footpatrol.com")
	dbot.CreateChannel("sns", "https://www.sneakersnstuff.com")
}

func checkKey(key string) bool {
	for _, usedKey := range usedKeys {
		if key == usedKey {
			return false
		}
	}
	return true
}

func ex(dbot *bot.Bot, pWrapper *pm.PatternsWrapper, productsWrapper *pm.ProductsWrapper) {
	log.Println("Terminaling...")
	pWrapper.ApplyStoreToDB()
	productsWrapper.ApplyStoreToDB()
	dbot.Session.Close()
}
