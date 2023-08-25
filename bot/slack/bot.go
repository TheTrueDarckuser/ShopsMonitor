package slack

import (
	"log"

	"github.com/nlopes/slack"
	pp "/restock-monitor/parser"
)

// Bot is a chat bot for slack
type Bot struct {
	BotToken        string
	Client          *slack.Client
	ProductsChannel <-chan *pp.Product
	Shop2Channels   map[string][]*slack.Channel
}

// NewBot creates a new slack bot
func NewBot(botToken string, productschannel <-chan *pp.Product) *Bot {
	bot := new(Bot)

	bot.BotToken = botToken
	bot.Shop2Channels = make(map[string][]*slack.Channel)
	bot.ProductsChannel = productschannel

	bot.Init()

	return bot
}

// Init initializes the bot
func (bot *Bot) Init() {
	bot.Client = slack.New(bot.BotToken)
}

// InitChannel creates a new channel
func (bot *Bot) InitChannel(channelName string, shopURLs ...string) (*slack.Channel, error) {
	channels, err := bot.Client.GetChannels(false)
	if err == nil {
		for _, channel := range channels {
			log.Println(channel.Name)
			if channel.Name == channelName {
				for _, shopURL := range shopURLs {
					bot.UpdateShopChannels(&channel, shopURL)
				}

				return &channel, err
			}
		}
	}
	return nil, err
}

// UpdateShopChannels updates shop2channels field
func (bot *Bot) UpdateShopChannels(channel *slack.Channel, shopURL string) {
	for _, t := range bot.Shop2Channels[shopURL] {
		if t.Name == channel.Name {
			return
		}
	}
	bot.Shop2Channels[shopURL] = append(bot.Shop2Channels[shopURL], channel)
}

// ListenProductsChannel listens not new arrivals channel and send them to fitting channel
func (bot *Bot) ListenProductsChannel() {
	for product := range bot.ProductsChannel {
		bot.SendProduct(product)
	}
}

// SendProduct send product to fitting channel
func (bot *Bot) SendProduct(product *pp.Product) error {
	for _, channel := range bot.Shop2Channels[product.ShopURL] {
		_, _, err := bot.SendMessageEmbed(product, *channel)
		if err != nil {
			// bot.SendProductOnFailure(channel.ID, product, 5, 3*time.Second)
			log.Printf("Erorr occured while sending %v: %s", product, err.Error())
			return err
		}
	}

	return nil
}

// SendMessageEmbed sends a message with an embedded product
func (bot *Bot) SendMessageEmbed(product *pp.Product, channel slack.Channel) (string, string, error) {
	return bot.Client.PostMessage(channel.ID, slack.MsgOptionText(
		product.Markdown(), false,
	))
}
