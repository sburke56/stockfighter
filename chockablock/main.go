package main

import (
	"flag"
	"fmt"
	"github.com/sburke56/stockfighter/support"
	"gopkg.in/sconf/ini.v0"
	"gopkg.in/sconf/sconf.v0"
	"time"
)

func main() {
	sconf.Must(&support.Cfg).Read(ini.File("config.gcfg"))
	fmt.Println(support.Cfg.Stockfighter.ApiKey)
	fmt.Println(support.Cfg.Stockfighter.Account)
	fmt.Println(support.Cfg.Stockfighter.Venue)
	fmt.Println(support.Cfg.Stockfighter.Symbol)
	fmt.Println(support.Cfg.Stockfighter.BaseUrl)

	lowPrice := flag.Int("low", 0, "price to prime engine with")
	flag.Parse()

	smallBlock := 5
	largeBlock := 1000

	openOrders := make(chan support.Order, 100)
	// The trick here I think is to peg the price of the stock
	// where you want so issue buys at the small block price lower
	// than the market, so there is demand for that price point &
	// the sellers come down to it.  When the price comes down low
	// enough issue the larger block buy.

	// This has to be done "slowly" and I adjusted the times
	// accordingly to space out the large block buys.
	for {
		quote, err := support.GetQuote(support.Cfg.Stockfighter.Venue, support.Cfg.Stockfighter.Symbol)
		if err == nil {
			if quote.Ask <= 0 {
				continue
			}

			fmt.Println(*lowPrice - quote.Ask)
			if (*lowPrice - quote.Ask) > 50 {
				support.Buy("limit", quote.Ask, largeBlock, openOrders)
				time.Sleep(2 * time.Second)
			}

			fmt.Println(quote)
			support.Buy("limit", quote.Ask-350, smallBlock, openOrders)
			time.Sleep(4 * time.Second)
		}
	}
}
