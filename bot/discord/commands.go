package discord

import (
	"html/template"
	"log"
	"reflect"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = template.FuncMap{
		"!help": help,
		//"!addproduct": addproduct,
	}
)

func commandsHandler(session *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, botprefix) && m.ID != botid {
		msg := strings.ToLower(strings.TrimSpace(m.Content))
		command := strings.Split(msg, " ")[0]

		fn, ok := commands[command]
		if !ok {
			_, err := session.ChannelMessageSend(m.ChannelID, "no such command")
			checkErr(err)
			return
		}
		v := reflect.ValueOf(fn)
		t := v.Type()

		if t.NumIn() == 0 {
			_, err := session.ChannelMessageSend(m.ChannelID, v.Call(nil)[0].String())
			checkErr(err)
			return
		}

		args := apply(strings.FieldsFunc(msg[len(command):], func(c rune) bool { return c == '"' }), strings.TrimSpace)
		args = concatArguments(args)
		if (len(args) != t.NumIn() && !t.IsVariadic()) || len(args) == 0 {
			_, err := session.ChannelMessageSend(m.ChannelID, "invalid number of arguments")
			checkErr(err)
			return
		}

		argv := make([]reflect.Value, len(args))
		for i := range argv {
			argv[i] = reflect.ValueOf(args[i])
		}
		_, err := session.ChannelMessageSend(m.ChannelID, v.Call(argv)[0].String())
		checkErr(err)
	}
}

func concatArguments(raw []string) []string {
	var args []string
	for _, s := range raw {
		if s != "" {
			args = append(args, s)
		}
	}
	return args
}

func apply(data []string, f func(string) string) []string {
	res := make([]string, len(data))
	for i, s := range data {
		res[i] = f(s)
	}
	return res
}

func checkErr(err error) {
	if err != nil {
		log.Printf("An error with bot: %v\n", err)
	}
}
