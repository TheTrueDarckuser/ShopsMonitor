package main

import (
	"fmt"
	"hash/fnv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type urlPair struct {
	Original string `bson:"original"`
	Short    string `bson:"short"`
}

type urlShortener struct {
	mongoSess *mgo.Session
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprint(h.Sum32())
}

// Short shortens a link
// returns true if links is new, false if it already existed
func (sh *urlShortener) Short(url string) (short string, new bool, err error) {
	sess := sh.mongoSess.Copy()

	// find pair from database
	var pair urlPair
	err = sess.DB("groc").C("urls").Find(bson.M{"original": url}).One(&pair)

	// if pair found
	if err == nil {
		return pair.Short, false, err
	}

	// if error has nothing to do with pair not existing
	if err != mgo.ErrNotFound {
		return "", false, err
	}

	short = hash(url)
	err = sess.DB("groc").C("urls").Insert(urlPair{
		Original: url,
		Short:    short,
	})
	if err != nil {
		return "", false, err
	}

	return "", true, nil
}

func (sh *urlShortener) Original(url string) (string, error) {
	var pair urlPair
	err := sh.mongoSess.DB("groc").C("urls").Find(bson.M{"short": url}).One(&pair)

	return pair.Short, err
}
