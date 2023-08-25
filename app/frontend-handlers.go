package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

func productsFrontHandler(w http.ResponseWriter, r *http.Request) {
	if err := template.Must(template.New("template").ParseGlob("templates/*")).ExecuteTemplate(w, "products.html", true); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addProductHandler(w http.ResponseWriter, r *http.Request) {
	if err := template.Must(template.New("template").ParseGlob("templates/*")).ExecuteTemplate(w, "addProduct.html", true); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	if err := template.Must(template.New("template").ParseGlob("templates/*")).ExecuteTemplate(w, "settings.html", true); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "session").(*sessions.Session)
	if v, ok := session.Values["access"]; ok && v.(bool) == true {
		http.Redirect(w, r, "/products", 302)
	}

	if err := template.Must(template.New("template").ParseGlob("templates/*")).ExecuteTemplate(w, "login.html", false); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func joinHandler(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "session").(*sessions.Session)
	if v, ok := session.Values["access"]; ok && v.(bool) == true {
		http.Redirect(w, r, "/products", 302)
	}

	if err := template.Must(template.New("template").ParseGlob("templates/*")).ExecuteTemplate(w, "join.html", false); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "session").(*sessions.Session)
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/auth", 302)
}

func discordHandler(w http.ResponseWriter, r *http.Request) {
	if err := template.Must(template.New("template").ParseGlob("templates/*")).ExecuteTemplate(w, "discord.html", true); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
