package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CreateChannel creates a new discord channel
func (b *Bot) CreateChannel(name string, shopURLs ...string) (channel *discordgo.Channel, err error) {
	channels, err := b.Session.GuildChannels(b.Guild.ID)
	if err != nil {
		return
	}

	name = strings.Replace(name, " ", "-", -1)
	if exists, chnl := checkChannelExists(name, channels); !exists {
		channel, err = b.Session.GuildChannelCreate(b.Guild.ID, name, "text")
		if err != nil {
			return nil, err
		}
	} else {
		channel = chnl
	}

	for _, shopURL := range shopURLs {
		b.UpdateShopChannels(channel, shopURL)
	}

	return
}

// checkChannelExists checks if channel already exists on server, if yes returns it
func checkChannelExists(name string, channels []*discordgo.Channel) (bool, *discordgo.Channel) {
	for _, chnl := range channels {
		if chnl.Name == name {
			return true, chnl
		}
	}
	return false, nil
}

// UpdateShopChannels updates channels assigned to a given shop
func (b *Bot) UpdateShopChannels(channel *discordgo.Channel, shopURL string) {
	for _, chnl := range b.Shop2Channels[shopURL] {
		if chnl.Name == channel.Name {
			return
		}
	}

	b.Shop2Channels[shopURL] = append(b.Shop2Channels[shopURL], channel)
}

// DeleteChannel delets channel
func (b *Bot) DeleteChannel(id string) {
	b.Session.ChannelDelete(id)
}
