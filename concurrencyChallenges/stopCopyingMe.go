// this task is about: sync and goroutine

package main

import (
	"flag"
	"log"
	"sync"
)

var messages = []string{
	"Hello!",
	"How are you?",
	"Are you just going to repeat what I say?",
	"So immature",
	"Stop copying me!",
}

// repeat Recursive concurrently prints out the given message n times
func repeatRecursive(n int, message string) {
	if n > 1 {
		log.Println(message)
		repeatRecursive(n-1, message)
	}
}

var wg sync.WaitGroup

// repeat concurrently prints out the given message n times
func repeat(n int, message string) {
	for i := 0; i < n; i++ {
		log.Println(message)
		wg.Done()
	}
}

// 1 rename to main
// 2 run in terminal: go run .\stopCopyingMe.go -factor 2
func main() {
	factor := flag.Int64("factor", 0, "The fan-out factor to repeat by")
	flag.Parse()
	n := int(*factor)
	for _, m := range messages {
		wg.Add(n)
		go repeat(n, m)
	}
	wg.Wait()
}
