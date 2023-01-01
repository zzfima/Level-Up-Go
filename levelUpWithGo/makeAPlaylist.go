// this task is about: sorting

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"
)

const path = "songs.json"

// Song stores all the song related information
type Song struct {
	Name      string `json:"name"`
	Album     string `json:"album"`
	PlayCount int64  `json:"play_count"`
}

// makePlaylist makes the merged sorted list of songs
func makePlaylist(albums [][]Song) (sortedSongs []Song) {
	for _, album := range albums {
		for _, song := range album {
			sortedSongs = append(sortedSongs, song)
		}
	}

	sort.SliceStable(sortedSongs, func(a, b int) bool {
		return sortedSongs[a].PlayCount > sortedSongs[b].PlayCount
	})
	return
}

// printTable prints merged playlist as a table
func printTable(songs []Song) {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "####\tSong\tAlbum\tPlay count")
	for i, s := range songs {
		fmt.Fprintf(w, "[%d]:\t%s\t%s\t%d\n", i+1, s.Name, s.Album, s.PlayCount)
	}
	w.Flush()

}

// importData reads the input data from file and creates the friends map
func importSongsEntities() (songsEntities [][]Song) {
	fileHandler, err := os.Open("songsEntities.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandler.Close()

	json.NewDecoder(fileHandler).Decode(&songsEntities)
	return
}

func mainMakeAPlaylist() {
	albums := importSongsEntities()
	playList := makePlaylist(albums)
	printTable(playList)
}
