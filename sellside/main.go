package main

import (
	"flag"
	"fmt"
	"gopkg.in/sconf/ini.v0"
	"gopkg.in/sconf/sconf.v0"
	"time"
	"github.com/sburke56/stockfighter/support"
	"sync/atomic"
)

func totalFills(fills []support.Fill) (total int){
	for _, f := range fills {
		total += f.Price*f.Qty
	}

	return total
}

func tallyExistingOrders(allOrders support.AllOrders, cash *int64, position *int64) {
	var position_ int64
	var cash_ int64

	for _, o := range allOrders.Orders {
		if !o.Open { // order is closed
			if o.Direction == "buy" {
				position_ += o.TotalFilled
				cash_ -= int64(totalFills(o.Fills))
			} else {
				position_ -= o.TotalFilled
				cash_ += int64(totalFills(o.Fills))
			}
		}
	}

	// store the cash & position values
	atomic.StoreInt64(cash, cash_)
	atomic.StoreInt64(position, position_)
}

func updateNav(pos *int64, cash *int64, nav *int64, done chan bool) {
	for {
		quote, _ := support.GetQuote()

		// get position * last value
		p := atomic.LoadInt64(pos)
		currentValue := p * quote.Last

		c := atomic.LoadInt64(cash)
		n := currentValue + c

		fmt.Printf("cash: %v  pos: %v  nav: %v\n", c, p, n)
		atomic.StoreInt64(nav, int64(n))
		time.Sleep(1 * time.Second)
	}

	done <- true
}

func getPosition(cash *int64, position *int64) {
	for {
		allOrders, err := support.GetStatusForStock()
		if err == nil {
			tallyExistingOrders(allOrders, cash, position)
		}

		time.Sleep(3 * time.Second)
	}
}

func buy(pricePt int64, spread int64, openOrders chan support.Order) {
	for {
		smallBlock := 5
		quote, _ := support.GetQuote()

		if (((pricePt - spread) < quote.Last) && (quote.Last < pricePt)) {
			support.Buy("limit", quote.Last, smallBlock, openOrders)
		}
		time.Sleep(300 * time.Millisecond)
	}
}

func sell(pricePt int64, spread int64, openOrders chan support.Order) {
	for {
		smallBlock := 5
		quote, _ := support.GetQuote()

		if ((pricePt < quote.Last) && (quote.Last < (pricePt + spread))) {
			support.Sell("limit", quote.Last, smallBlock, openOrders)
		}
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	sconf.Must(&support.Cfg).Read(ini.File("config.gcfg"))
	fmt.Println(support.Cfg.Stockfighter.ApiKey)
	fmt.Println(support.Cfg.Stockfighter.Account)
	fmt.Println(support.Cfg.Stockfighter.Venue)
	fmt.Println(support.Cfg.Stockfighter.Symbol)
	fmt.Println(support.Cfg.Stockfighter.BaseUrl)

	pricePt := flag.Int64("price", 0, "price point to start at")
	spread := flag.Int64("spread", 0, "spread to buy/sell at")
	flag.Parse()

	var cash int64
	var position int64
	var nav int64

	openOrders := make(chan support.Order, 100)
	done := make(chan bool)

	go getPosition(&cash, &position)
	go updateNav(&position, &cash, &nav, done)

	go buy(*pricePt, *spread, openOrders)
	go sell(*pricePt, *spread, openOrders)

	<- done
}
