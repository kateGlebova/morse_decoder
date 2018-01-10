package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const DatabaseName = "morse"
const CollectionName = "codes"

type Code struct {
	Encoded string `bson: "encoded" json: "encoded"`
	Decoded string `bson: "decoded" json: "decoded"`
}

func (code Code) String() string {
	return fmt.Sprintf("'%s':'%s'", code.Encoded, code.Decoded)
}

type MorseCode struct {
	collection *mgo.Collection
	cache      *cache.Cache
}

func NewMorseCode(session *mgo.Session) *MorseCode {
	codesCollection := session.DB(DatabaseName).C(CollectionName)
	codesCache := cache.New(5*time.Minute, 10*time.Minute)
	m := MorseCode{codesCollection, codesCache}
	return &m
}

func (this MorseCode) get_decoded(encoded string) (string, error) {
	result := Code{}
	decoded, found := this.cache.Get(encoded)
	if found {
		return decoded.(string), nil
	} else {
		err := this.collection.Find(bson.M{"encoded": encoded}).One(&result)
		if err != nil {
			return "", errors.New(fmt.Sprintf("'%s' is not a valid Morse code.", encoded))
		} else {
			this.cache.Set(encoded, result.Decoded, cache.DefaultExpiration)
			return result.Decoded, nil
		}
	}
}

func (this MorseCode) Decode(input string) string {
	buf := bytes.NewBufferString("")
	for _, c := range strings.Split(input, " ") {
		symbol, err := this.get_decoded(c)
		if err != nil {
			continue
		}
		buf.WriteString(symbol)
	}
	return strings.TrimSpace(buf.String())
}
