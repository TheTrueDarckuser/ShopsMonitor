package slack

import (
	"testing"

	"github.com/nlopes/slack"
	"/restock-monitor/parser"
)

const token = "xoxb-386832718983-524342788519-Y14JZOvZtpeHXvuNTGoMMEV9"

var (
	productChannel = make(chan *parser.Product)
	slackBot       = NewBot(token, productChannel)
	api            = slack.New(token)
)

func TestInitChannel(t *testing.T) {
	chanName := "testChannel"

	channel, err := slackBot.InitChannel(chanName, "https://penis.com")
	if err != nil {
		t.Error(err)
	}

	_ = channel
}

func TestSendProduct(t *testing.T) {
	err := slackBot.SendProduct(&parser.Product{
		Name: "A Test Product",
		Link: "https://goatse.com",
	})

	if err != nil {
		t.Error(err)
	}
}
