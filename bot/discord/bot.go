package discord

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	pm "/restock-monitor/monitor-lff"
	pp "/restock-monitor/parser"
)

// Bot is a chat bot for Discord
type Bot struct {
	BotToken       string
	Session        *discordgo.Session
	Guild          *discordgo.Guild
	botUser        *discordgo.User
	newlyAvailable <-chan *pm.ProductWrapper
	Shop2Channels  map[string][]*discordgo.Channel
}

func botReady(s *discordgo.Session, m *discordgo.Ready) {
	s.UpdateStatus(0, "Grocodyle inc.|!help")
}

// NewBot creates a new DiscorBot instance
func NewBot(botToken, guildID string, newlyAvailable <-chan *pm.ProductWrapper) *Bot {
	bot := new(Bot)

	bot.BotToken = botToken
	bot.Shop2Channels = make(map[string][]*discordgo.Channel)
	bot.newlyAvailable = newlyAvailable

	err := bot.Init(guildID, botReady, commandsHandler)
	if err != nil {
		panic(err)
	}

	return bot
}

// Init initializes bot session
func (b *Bot) Init(guildID string, handlers ...interface{}) (err error) {
	b.Session, err = discordgo.New("Bot " + b.BotToken)
	if err != nil {
		return
	}

	b.botUser, err = b.Session.User("@me")
	if err != nil {
		return
	}

	b.Guild, err = b.Session.Guild(guildID)
	if err != nil {
		return
	}

	for _, hander := range handlers {
		b.Session.AddHandler(hander)
	}

	b.Session.Client.Timeout = 10 * time.Second
	b.Session.SyncEvents = false

	return b.Session.Open()
}

// ListenProductsChannel listens to products channel
func (b *Bot) ListenProductsChannel() {
	go func() {
		for {
			select {
			case v := <-b.newlyAvailable:
				b.SendProduct(v.Product)
				v.Type = "restock"
			}
		}
	}()
}

// SendProduct sends product to channels
func (b *Bot) SendProduct(product *pp.Product) {
	for _, channel := range b.Shop2Channels[product.ShopURL] {
		_, err := b.SendProductEmbed(product, channel.ID)
		if err != nil {
			b.SendProductOnFailure(channel.ID, product, 5, 3*time.Second)
			log.Printf("Erorr occured while sending %v: %s", product, err.Error())
		}
	}
}

// SendProductOnFailure sends product to channels
func (b *Bot) SendProductOnFailure(channelID string, product *pp.Product, times int, delay time.Duration) {
	for i := 0; i < times; i++ {
		_, err := b.SendProductEmbed(product, channelID)
		if err != nil {
			log.Printf("Erorr occured while retrying to send %v: %s", product, err.Error())
		} else {
			return
		}
		time.Sleep(delay)
	}
}

// SendProductEmbed send an embedded message to channel containing product info
func (b *Bot) SendProductEmbed(product *pp.Product, channelID string) (*discordgo.Message, error) {
	if product.Type == "coming soon" {
		return b.Session.ChannelMessageSendEmbed(channelID, getComingSoonEmbed(product))
	}
	if atc := embedAtc(product); atc != nil {
		return b.Session.ChannelMessageSendEmbed(channelID, getFullEmbed(product, atc))
	}
	return nil, nil
}
