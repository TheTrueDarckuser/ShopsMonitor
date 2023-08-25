package discord

import (
	"testing"

	pp "/restock-monitor/parser"
)

func TestParseDrugUrl(t *testing.T) {
	product := pp.Product{
		Name:    "Name",
		Link:    "https://www.solebox.com/en/Footwear/Basketball/Air-Force-1-Skeleton-QS.html",
		ShopURL: "https://google.com",

		PhotoURL: "https://gitlab.com/uploads/-/system/user/avatar/1818202/avatar.png?width=40",

		Type: "nibba",
	}

	bot := NewBot("NDk2MDcxMjI5OTE3MTAyMTAx.DpbFug.oOXqeOtrD-HlzKPotyZfosbOmgI", "413763893043789824", make(<-chan *pp.Product))

	_, err := bot.SendProductEmbed(&product, "500555848490156044")
	if err != nil {
		t.Error(err)
	}
}
