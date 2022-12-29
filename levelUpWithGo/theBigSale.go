// this task is about: filter and sort a slice

package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sort"
)

// SaleItem represents the item part of the big sale.
type SaleItem struct {
	Name           string  `json:"name"`
	OriginalPrice  float64 `json:"originalPrice"`
	ReducedPrice   float64 `json:"reducedPrice"`
	SalePercentage float64
}

// matchSales adds the sales procentage of the item
// and sorts the array accordingly.
func matchSales(budget float64, items []SaleItem) []SaleItem {
	var saleItems []SaleItem
	for _, item := range items {
		if item.ReducedPrice <= budget {
			item.SalePercentage = 100.0 - 100.0*item.ReducedPrice/item.OriginalPrice
			saleItems = append(saleItems, item)
		}
	}
	sort.SliceStable(saleItems,
		func(p, q int) bool {
			return saleItems[p].SalePercentage < saleItems[q].SalePercentage
		})

	return saleItems
}

// printItems prints the items and their sales.
func printItems(items []SaleItem) {
	log.Println("The BIG sale has started with our amazing offers!")
	if len(items) == 0 {
		log.Println("No items found.:( Try increasing your budget.")
	}
	for i, r := range items {
		log.Printf("[%d]:%s is %.2f OFF! Get it now for JUST %.2f!\n",
			i, r.Name, r.SalePercentage, r.ReducedPrice)
	}
}

// importData reads the raffle entries from file and
// creates the entries slice.
func importSalesData() []SaleItem {
	fileHandler, err := os.Open("salesEntities.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandler.Close()

	var salesItems []SaleItem
	decoder := json.NewDecoder(fileHandler)
	decoder.Decode(&salesItems)

	return salesItems
}

// 1 rename to main
// 2 run in terminal: go run .\theBigSale.go -budget 500
func mainTheBigSale() {
	budget := flag.Float64("budget", 0.0, "The max budget you want to shop with.")
	flag.Parse()
	salesItems := importSalesData()
	matchedItems := matchSales(*budget, salesItems)
	printItems(matchedItems)
}
