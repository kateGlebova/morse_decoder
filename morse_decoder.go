package morse_decoder

import(
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"bytes"
	"strings"
)

const DatabaseName = "morse"
const CollectionName = "codes"

type Code struct {
	Encoded string 
	Decoded string 
}

func (code Code) String() string {
	return fmt.Sprintf("'%s':'%s'", code.Encoded, code.Decoded)
}

type MorseCode struct {
	collection *mgo.Collection
}

func NewMorseCode(session *mgo.Session) *MorseCode {
	m := MorseCode{session.DB(DatabaseName).C(CollectionName)}
	return &m
}
	
func (this MorseCode) get_decoded(encoded string) (string, error) {
	result := Code{}
	err := this.collection.Find(bson.M{"encoded": encoded}).One(&result)
    if err != nil {
        return "", errors.New(fmt.Sprintf("'%s' is not a valid Morse code.", encoded))
    } else {
        return result.Decoded, nil
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
