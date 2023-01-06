// this task is about: chanel and wait group

package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const maxSeconds = 1

var wg1 sync.WaitGroup
var chOwner1 = make(chan struct{})
var chOwner2 = make(chan struct{})

var chDog1 = make(chan struct{})
var chDog2 = make(chan struct{})

type dog struct {
	name string
}

type owner struct {
	name string
}

func (d dog) fetchLeash() {
	log.Printf("%s goes to fetch leash.\n", d.name)
	randomSleep()
	log.Printf("%s has fetched leash. Woof woof!\n", d.name)
	chDog1 <- struct{}{}
	wg1.Done()
}

func (d dog) findTreats() {
	<-chDog1
	log.Printf("%s goes to fetch treats.\n", d.name)
	randomSleep()
	log.Printf("%s has fetched the treats. Woof woof!\n", d.name)
	chDog2 <- struct{}{}
	wg1.Done()
}

func (d dog) runOutside() {
	<-chDog2
	log.Printf("%s starts running outside.\n", d.name)
	randomSleep()
	log.Printf("%s is having fun outside. Woof woof!\n", d.name)
	wg1.Done()
}

func (o owner) putShoesOn() {
	log.Printf("%s starts putting shoes on.\n", o.name)
	randomSleep()
	log.Printf("%s finishes putting shoes on.\n", o.name)
	chOwner1 <- struct{}{}
	wg1.Done()
}

func (o owner) findKeys() {
	<-chOwner1
	log.Printf("%s starts looking for keys.\n", o.name)
	randomSleep()
	log.Printf("%s has found keys.\n", o.name)
	chOwner2 <- struct{}{}
	wg1.Done()
}

func (o owner) lockDoor() {
	<-chOwner2
	log.Printf("%s starts locking the door.\n", o.name)
	randomSleep()
	log.Printf("%s has locked the door.\n", o.name)
	wg1.Done()
}

func randomSleep() {
	r := rand.Intn(maxSeconds)
	time.Sleep(time.Duration(r)*time.Second + 500*time.Millisecond)
}

func executeWalk(ownerActions []func(), dogActions []func()) {
	for _, o := range ownerActions {
		go o()
	}

	for _, d := range dogActions {
		go d()
	}
}

func mainWalkTheDog() {
	wg1.Add(6)
	defer wg1.Wait()

	owner := owner{name: "Jimmy"}
	dog := dog{name: "Lucky"}
	ownerActions := []func(){
		owner.putShoesOn,
		owner.findKeys,
		owner.lockDoor,
	}

	dogActions := []func(){
		dog.fetchLeash,
		dog.findTreats,
		dog.runOutside,
	}
	executeWalk(ownerActions, dogActions)
}
