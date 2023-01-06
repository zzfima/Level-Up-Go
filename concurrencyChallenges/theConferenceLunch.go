// this task is about: chanel and wait group

package main

import (
	"log"
	"sync"
)

// the number of attendees we need to serve lunch to
const consumerCount = 300

var ch = make(chan struct{})
var wg2 sync.WaitGroup

// foodCourses represents the types of resources to pass to the consumers
var foodCourses = []string{
	"Caprese Salad",
	"Spaghetti Carbonara",
	"Vanilla Panna Cotta",
}

// takeLunch is the consumer function for the lunch simulation
// Change the signature of this function as required
func takeLunch(name string) {
	<-ch
	log.Printf("consume " + name + " ....")
	wg2.Done()
}

// serveLunch is the producer function for the lunch simulation.
// Change the signature of this function as required
func serveLunch(course string) {
	log.Printf("make " + course + " ....")
	ch <- struct{}{}
}

func mainTheConferenceLunch() {
	log.Printf("Welcome to the conference lunch! Serving %d attendees.\n", consumerCount)
	for i := 0; i < len(foodCourses); i++ {
		wg2.Add(1)
		go serveLunch(foodCourses[i])
		go takeLunch(foodCourses[i])
	}

	wg2.Wait()
}
