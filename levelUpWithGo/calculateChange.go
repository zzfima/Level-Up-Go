// this task is about: algorithmic problem, exploring maps and slices

package main

import (
	"flag"
	"log"
	"math"
)

// coin contains the name and value of a coin
type coin struct {
	name  string
	value float64
}

// coins is the list of values available for making change.
var coins = []coin{
	{name: "1 pound", value: 1},
	{name: "50 pence", value: 0.50},
	{name: "20 pence", value: 0.20},
	{name: "10 pence", value: 0.10},
	{name: "5 pence", value: 0.05},
	{name: "1 penny", value: 0.01},
}

// calculateChange returns the coins required to calculate the
func calculateChange(amount float64) (coinsMap map[coin]int) {

	coinsMap = make(map[coin]int)

	for _, cn := range coins {
		if amount >= cn.value {
			count := math.Floor(amount / cn.value)
			amount -= count * cn.value
			coinsMap[cn] = int(count)
		}
	}
	return
}

// printCoins prints all the coins in the slice to the terminal.
func printCoins(change map[coin]int) {
	if len(change) == 0 {
		log.Println("No change found.")
		return
	}
	log.Println("Change has been calculated.")
	for coin, count := range change {
		log.Printf("%d x %s \n", count, coin.name)
	}
}

// 1 rename to main
// 2 run in terminal: go run .\calculateChange.go -amount 34.30
func mainCalculateChange() {
	amount := flag.Float64("amount", 0.0, "The amount you want to make change for")
	flag.Parse()
	if *amount == 0.0 {
		log.Println("\nCommand line arguments:")
		flag.PrintDefaults()
		return
	}

	change := calculateChange(*amount)
	printCoins(change)
}
