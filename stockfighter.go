package main

import (
	"fmt"
	"time"
	"flag"
)

const ApiKey = "dfb3fcc6403b19e213743bebedb43232b7623b00"
const Account = "MFB38715308"
const Symbol = "OMEU"
const Venue = "SLTEX"
const BaseUrl = "https://api.stockfighter.io/ob/api"

func main() {
	lowPrice := flag.Int("low", 0, "price to prime engine with")
	flag.Parse()

	smallBlock := 5
	largeBlock := 1000

	// The trick here I think is to peg the price of the stock
	// where you want so issue buys at the small block price lower
	// than the market, so there is demand for that price point &
	// the sellers come down to it.  When the price comes down low
	// enough issue the larger block buy.

	// This has to be done "slowly" and I adjusted the times
	// accordingly to space out the large block buys.
	for {
		quote, err := getQuote(Venue, Symbol)
		if err == nil {
			if quote.Ask <= 0 {
				continue
			}

			fmt.Println(*lowPrice - quote.Ask)
			if ((*lowPrice - quote.Ask) > 50) {
				buy("limit", quote.Ask, largeBlock)
				time.Sleep(2*time.Second)
			}

			fmt.Println(quote)
			buy("limit", quote.Ask-350, smallBlock)
			time.Sleep(4*time.Second)
		}
	}
}
