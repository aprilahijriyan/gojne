package main

import (
	"encoding/json"
	"fmt"

	"github.com/aprilahijriyan/gojne"
)

func main() {
	// Scraped data
	var trackingNumber = "3452440340005"
	// get tracking information
	data := gojne.GetTracking(trackingNumber)
	// Output
	a, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(a))
}
