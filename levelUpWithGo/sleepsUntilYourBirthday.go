package main

import (
	"fmt"
	"log"
	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	t, _ := time.Parse(expectedFormat, target)
	return t
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	today := time.Now()
	d := time.Until(target)
	fmt.Println(today, d)
	return d.Hours() / 24
}

func main() {
	bday := "2023-06-21"
	target := parseTime(bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}
