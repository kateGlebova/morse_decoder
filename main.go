package main

import (
	"bufio"
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

func main() {
	db, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("Cannot dial MongoDB", err)
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)
	var storageString string
	fmt.Print("Enter Morse encoded string:")
	storageString, _ = reader.ReadString('\n')

	morse_code := NewMorseCode(db)
	decoded := morse_code.Decode(storageString[:len(storageString)-1])

	fmt.Println(decoded)
}
