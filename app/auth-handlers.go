package main

import (
	"crypto/sha256"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//		########
//		# User #
//		########

// User is an abstract monitor user
type User struct {
	Username     string   `bson:"username"`
	PasswordHash [32]byte `bson:"phash"`
}

// NewUser creates a new User instance
func NewUser(uname string, password string) *User {
	user := new(User)

	user.Username = uname
	user.PasswordHash = sha256.Sum256([]byte(password))

	return user
}

// validPassword validates passwords
func (u User) validPassword(password string) bool {
	return sha256.Sum256([]byte(password)) == u.PasswordHash
}

//		##############
//		# InviteKeys #
//		##############

// InviteKey is used to register users
type InviteKey struct {
	Value   string    `json:"value" bson:"value"`
	Expires time.Time `json:"expires" bson:"expires"`
}

// On access denied handlers

// authRedirect redirects user to login page
func authRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/auth", 302)
}

// accessDenied writes Forbidden status to http writer
func accessDenied(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Access denied", http.StatusForbidden)
}

// grantAccessHandler grants access to user
func grantAccessHandler(w http.ResponseWriter, r *http.Request, session *sessions.Session) {
	session.Values["access"] = true
	session.Save(r, w)
}

// handleAuth authenficates a user
func handleAuth(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := context.Get(r, "db").(*mgo.Database)

	var user User
	err = db.C("users").Find(bson.M{"username": request.Username}).One(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.validPassword(request.Password) {
		grantAccessHandler(w, r, context.Get(r, "session").(*sessions.Session))
	} else {
		http.Error(w, "Something's invalid", http.StatusForbidden)
		return
	}
}

// handleJoin joins a user
func handleJoin(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username  string
		Password  string
		InviteKey string
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := context.Get(r, "db").(*mgo.Database)
	if count, err := db.C("users").Find(bson.M{"username": request.Username}).Count(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if count > 0 {
		http.Error(w, "Username already taken", http.StatusIMUsed)
		return
	}

	var key InviteKey
	if err := db.C("keys").Find(bson.M{"value": request.InviteKey}).One(key); err != nil {
		http.Error(w, "Key's not valid", http.StatusForbidden)
		return
	}

	if err := db.C("keys").Remove(key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := db.C("users").Insert(NewUser(request.Username, request.Password)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grantAccessHandler(w, r, context.Get(r, "session").(*sessions.Session))
}
