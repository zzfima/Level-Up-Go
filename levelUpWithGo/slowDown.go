// this task is about: strings manipulation
package main

import (
	"flag"
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	words := strings.Split(msg, " ")
	for _, w := range words {
		var pw []string
		for index, c := range w {
			rb := strings.Repeat(string(c), index+1)
			pw = append(pw, rb)
		}
		print(strings.Join(pw, ""))
	}
}

// 1 rename to main
// 2 run in terminal: go run .\slowDown.go -msg "helko pupsik"
func mainSlowDown() {
	msg := flag.String("msg", "", "Sentence to process")
	flag.Parse()
	if *msg == "" {
		flag.PrintDefaults()
		return
	}

	slowDown(*msg)
}
