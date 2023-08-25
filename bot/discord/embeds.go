package discord

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/bwmarrin/discordgo"
	pp "/restock-monitor/parser"
)

func getComingSoonEmbed(product *pp.Product) *discordgo.MessageEmbed {
	t := pp.Global[product.ShopURL]
	return &discordgo.MessageEmbed{
		Color: 0xff0000,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "**Type**:",
				Value:  "**COMING SOON**",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Product**:",
				Value:  fmt.Sprintf("**[%s](%s)**", strings.ToUpper(product.Name), product.Link),
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Shop**:",
				Value:  fmt.Sprintf("**%s**", t.ShopName),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Basket**:",
				Value:  fmt.Sprintf("**%s%s**", product.ShopURL, t.Basket),
				Inline: true,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: product.PhotoURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Grocodile",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func getFullEmbed(product *pp.Product, atc *discordgo.MessageEmbedField) *discordgo.MessageEmbed {
	t := pp.Global[product.ShopURL]
	return &discordgo.MessageEmbed{
		Color: 0x3B6BCC,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "**Product**:",
				Value:  fmt.Sprintf("**[%s](%s)**", strings.ToUpper(product.Name), product.Link),
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Type**:",
				Value:  fmt.Sprintf("**%s**", product.Type),
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Shop**:",
				Value:  fmt.Sprintf("**%s**", t.ShopName),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Basket**:",
				Value:  fmt.Sprintf("**%s%s**", product.ShopURL, t.Basket),
				Inline: true,
			},
			atc,
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: product.PhotoURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Grocodile",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func embedAtc(p *pp.Product) *discordgo.MessageEmbedField {
	if v, ok := pp.AtcM[p.ShopURL]; ok {
		if atc := v(p.Link); len(atc) > 0 {
			return &discordgo.MessageEmbedField{
				Name:   "**ATC**:",
				Value:  getStringFromMap(atc),
				Inline: false,
			}
		}
		log.Printf("Error occured while retrieving %s atc for %s", p.ShopURL, p.Link)
		return nil
	}
	return &discordgo.MessageEmbedField{
		Name:   "**ATC**:",
		Value:  "**no atc**",
		Inline: false,
	}
}

func getStringFromMap(m map[string]string) string {
	sizes := false
	for _, v := range m {
		if v == "" {
			sizes = true
			break
		}
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Sort(bySize(keys))

	var s string
	if sizes {
		for _, key := range keys {
			s = s + fmt.Sprintf("**%s**\n", key)
		}
	} else {
		for _, key := range keys {
			s = s + fmt.Sprintf("[%s](%s)\n", key, m[key])
		}
	}

	return s
}

// --------------------------- Sorting --------------------------- //

type bySize []string

func (s bySize) Len() int {
	return len(s)
}

func (s bySize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s bySize) Less(i, j int) bool {
	a, err := getNum(s[i])
	if err != nil {
		return true
	}
	b, _ := getNum(s[j])
	return a < b
}

func getNum(s string) (int, error) {
	i := 0
	var t string
	for i < len(s) {
		if unicode.IsDigit(rune(s[i])) {
			break
		}
		i++
	}

	for i < len(s) {
		if !unicode.IsDigit(rune(s[i])) {
			break
		} else {
			t = t + string(s[i])
		}
		i++
	}

	return strconv.Atoi(t)
}
