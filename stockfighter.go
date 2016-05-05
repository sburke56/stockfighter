package main

import (
	"flag"
	"fmt"
	"gopkg.in/sconf/ini.v0"
	"gopkg.in/sconf/sconf.v0"
	"time"
)

var Cfg Config

func main() {
	sconf.Must(&Cfg).Read(ini.File("config.gcfg"))
	fmt.Println(Cfg.Stockfighter.ApiKey)
	fmt.Println(Cfg.Stockfighter.Account)
	fmt.Println(Cfg.Stockfighter.Venue)
	fmt.Println(Cfg.Stockfighter.Symbol)
	fmt.Println(Cfg.Stockfighter.BaseUrl)

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
		quote, err := getQuote(Cfg.Stockfighter.Venue, Cfg.Stockfighter.Symbol)
		if err == nil {
			if quote.Ask <= 0 {
				continue
			}

			fmt.Println(*lowPrice - quote.Ask)
			if (*lowPrice - quote.Ask) > 50 {
				buy("limit", quote.Ask, largeBlock)
				time.Sleep(2 * time.Second)
			}

			fmt.Println(quote)
			buy("limit", quote.Ask-350, smallBlock)
			time.Sleep(4 * time.Second)
		}
	}
}
