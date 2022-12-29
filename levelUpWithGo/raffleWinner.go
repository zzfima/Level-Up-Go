// this task is about: JSON processing and random number generation knowledge.

package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// importData reads the raffle entries from file and creates the entries slice.
func importUsersData() []raffleEntry {
	entriesFileHandler, e := os.Open("usersEntries1.json")
	if e != nil {
		log.Fatal("Can not open entries file")
	}
	defer entriesFileHandler.Close()
	decoder := json.NewDecoder(entriesFileHandler)
	var deserializedArray []raffleEntry
	decoder.Decode(&deserializedArray)

	return deserializedArray
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func mainRaffleWinner() {
	entries := importUsersData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}
