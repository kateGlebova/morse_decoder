# Morse Decoder

An exercise done within NTUU "KPI" Computer Architecture course in 3 stages:
1. Implement a simple Morse Decoder

Morse code is stored in MongoDB. Name of the database and collection are declared as constants in morse_decoder.go file.

To run the current code:
* Run a mongo container
```bash
$sudo docker run --name mongo-db -p 27017:27017 -d mongo
```
* Open the mongo shell
```bash
$sudo docker exec -it mongo-db mongo
```
* Put morse code in collection 'codes' in database 'morse'
```
> use morse
> db.codes.insertMany([
    {'encoded': '-----', 'decoded': '0'},
    {'encoded': '.----', 'decoded': '1'},
    {'encoded': '..---', 'decoded': '2'},
    {'encoded': '...--', 'decoded': '3'},
    {'encoded': '....-', 'decoded': '4'},
    {'encoded': '.....', 'decoded': '5'},
    {'encoded': '-....', 'decoded': '6'},
    {'encoded': '--...', 'decoded': '7'},
    {'encoded': '---..', 'decoded': '8'},
    {'encoded': '----.', 'decoded': '9'},
    {'encoded': '.-', 'decoded': 'A'},
    {'encoded': '-...', 'decoded': 'B'},
    {'encoded': '-.-.', 'decoded': 'C'},
    {'encoded': '-..', 'decoded': 'D'},
    {'encoded': '.', 'decoded': 'E'},
    {'encoded': '..-.', 'decoded': 'F'},
    {'encoded': '--.', 'decoded': 'G'},
    {'encoded': '....', 'decoded': 'H'},
    {'encoded': '..', 'decoded': 'I'},
    {'encoded': '.---', 'decoded': 'J'},
    {'encoded': '-.-', 'decoded': 'K'},
    {'encoded': '.-..', 'decoded': 'L'},
    {'encoded': '--', 'decoded': 'M'},
    {'encoded': '-.', 'decoded': 'N'},
    {'encoded': '---', 'decoded': 'O'},
    {'encoded': '.--.', 'decoded': 'P'},
    {'encoded': '--.-', 'decoded': 'Q'},
    {'encoded': '.-.', 'decoded': 'R'},
    {'encoded': '...', 'decoded': 'S'},
    {'encoded': '-', 'decoded': 'T'},
    {'encoded': '..-', 'decoded': 'U'},
    {'encoded': '...-', 'decoded': 'V'},
    {'encoded': '.--', 'decoded': 'W'},
    {'encoded': '-..-', 'decoded': 'X'},
    {'encoded': '-.--', 'decoded': 'Y'},
    {'encoded': '--..', 'decoded': 'Z'},
    {'encoded': '.-.-.-', 'decoded': '.'},
    {'encoded': '--..--', 'decoded': ','},
    {'encoded': '..--..', 'decoded': '?'},
    {'encoded': '-.-.--', 'decoded': '!'},
    {'encoded': '-....-', 'decoded': '-'},
    {'encoded': '-..-.', 'decoded': '/'},
    {'encoded': '.--.-.', 'decoded': '@'},
    {'encoded': '-.--.', 'decoded': '('},
    {'encoded': '-.--.-', 'decoded': ')'}
  ])
```
* Install dependencies
```bash
$glide install
```
* Compile the package
```bash
$go install
```
* Run tests
```bash
$go test
```
