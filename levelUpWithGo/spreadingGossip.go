// this task is about: recursion

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// Friend represents a friend and their connections.
type Friend struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Friends []string `json:"friends"`
}

// hearGossip indicates that the friend has heard the gossip.
func (f *Friend) hearGossip() {
	log.Printf("%s has heard the gossip!\n", f.Name)
}

// Friends represents the map of friends and connections
type Friends struct {
	friendsMap map[string]Friend
}

// getFriend fetches the friend given an id.
func (f *Friends) getFriend(id string) Friend {
	return f.friendsMap[id]
}

// getRandomFriend returns an random friend.
func (f *Friends) getRandomFriend() Friend {
	rand.Seed(time.Now().Unix())
	id := (rand.Intn(len(f.friendsMap)-1) + 1) * 100
	return f.getFriend(fmt.Sprint(id))
}

var gossipMap = make(map[string]int)

// spreadGossip ensures that all the friends in the map have heard the news
func spreadGossip(root Friend, friends Friends) {
	for _, friendId := range root.Friends {
		if gossipMap[friendId] == 0 {
			gossipMap[friendId]++
			f := friends.getFriend(friendId)
			f.hearGossip()
			spreadGossip(f, friends)
		}
	}
}

// importData reads the input data from file and
// creates the friends map.
func importFriendsEntities() Friends {
	fileHandler, err := os.Open("friendsEntities.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandler.Close()

	var friendsEntities []Friend
	decoder := json.NewDecoder(fileHandler)
	decoder.Decode(&friendsEntities)

	friendsEntitiesMap := make(map[string]Friend, len(friendsEntities))
	for _, d := range friendsEntities {
		friendsEntitiesMap[d.ID] = d
	}

	return Friends{
		friendsMap: friendsEntitiesMap,
	}
}

func main() {
	friends := importFriendsEntities()
	root := friends.getRandomFriend()
	root.hearGossip()
	spreadGossip(root, friends)
}
