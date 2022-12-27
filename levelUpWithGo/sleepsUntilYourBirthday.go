package main

import (
	"flag"
	"log"
	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) (parsedTime time.Time) {
	parsedTime, e := time.Parse(expectedFormat, target)
	isPastDate := time.Now().After(parsedTime)

	if e != nil || isPastDate {
		log.Fatal("invalid date")
	}
	return
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	d := time.Until(target)
	return d.Hours() / 24
}

// 1. rename to main
// 2. run in terminal: go run .\sleepsUntilYourBirthday.go -bday 2023-06-21
func mainSleepsUntilYourBirthday() {
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!", int(calcSleeps(target)))
}
