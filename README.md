# Morse Decoder

An exercise done within NTUU "KPI" Computer Architecture course in 3 stages:
1. Implement a simple Morse Decoder

Morse code is stored in MongoDB. Name of the database and collection are declared as constants in morse_decoder.go file. 

To run the current code:
* Run a mongo container
```bash
sudo docker run --name mongo-db -p 27017:27017 -d mongo
```
* Open the mongo shell
```bash
sudo docker exec -it mongo-db mongo
```
* Put morse code in collection 'codes' in database 'morse'
```
> use morse
> db.codes.insertMany([])
```