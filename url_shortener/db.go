package main

import (
	"crypto/sha1"
	"errors"
	"fmt"

	"github.com/go-martini/martini"
)

type Link struct {
	Url    string `form:"url"`
	Visits int
}

func (link Link) Hash() string {

	hash := sha1.New()
	hash.Write([]byte(link.Url))
	sha := fmt.Sprintf("%x", hash.Sum(nil))

	var small = ""

	for i := 0; i < 10; i++ {

		small += fmt.Sprintf("%c", sha[i])
	}

	return small
}

type linkdb map[string]Link

func (db linkdb) Get(key string) (Link, error) {

	link, ok := db[key]

	if ok {

		return link, nil

	} else {

		return Link{}, errors.New("not found")
	}
}

func (db linkdb) Save(link Link) (string, error) {

	var hash = link.Hash()

	_, err := db.Get(hash)

	if err != nil {

		db[hash] = link
	}

	return hash, nil
}

func (db linkdb) SaveVisit(link Link) error {

	var hash = link.Hash()

	link, err := db.Get(hash)

	if err == nil {

		link.Visits++
		db[hash] = link

		return nil
	}

	return errors.New("Not existing link")
}

// Martini mapper

var database linkdb

func DB() martini.Handler {

	//Initialize DB
	database = linkdb{}

	return func(c martini.Context) {

		c.Map(database)
		c.Next()
	}
}
