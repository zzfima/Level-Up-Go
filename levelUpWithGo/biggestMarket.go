// this task is about: using maps

package main

import (
	"encoding/json"
	"log"
	"os"
)

// User represents a user record.
type User struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

// getBiggestMarket takes in the slice of users and
// returns the biggest market.
func getBiggestMarket(users []User) (string, int) {
	countriesCount := make(map[string]int)
	for _, user := range users {
		countriesCount[user.Country]++
	}

	countryName := ""
	usersAmount := 0

	for n, v := range countriesCount {
		if v > usersAmount {
			usersAmount = v
			countryName = n
		}
	}

	return countryName, usersAmount
}

// importData reads the raffle entries from file and
// creates the entries slice.
func importData() []User {
	fileHandler, err := os.Open("usersEntities2.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandler.Close()

	var users []User
	decoder := json.NewDecoder(fileHandler)
	decoder.Decode(&users)

	return users
}

func mainBiggestMarket() {
	users := importData()
	country, count := getBiggestMarket(users)
	log.Printf("The biggest user market is %s with %d users.\n", country, count)
}
