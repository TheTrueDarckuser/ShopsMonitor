package main

import (
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

var formDecoder = schema.NewDecoder()

func oauthCallback(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()

		var response struct {
			AccessToken  string `schema:"access_token"`
			TokenType    string `schema:"token_type"`
			ExpiresIn    string `schema:"expires_in"`
			RefreshToken string `schema:"refresh_token"`
			Scope        string `schema:"scope"`
		}

		if err := formDecoder.Decode(&response, r.PostForm); err == nil {
			log.Println(response)
		} else {
			log.Println(err)
		}
	}
}

func redirectDiscordOauth() {

}
