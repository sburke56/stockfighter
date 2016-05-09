package main

import (
//	"flag"
	"fmt"
	"gopkg.in/sconf/ini.v0"
	"gopkg.in/sconf/sconf.v0"
	"time"
	"github.com/sburke56/stockfighter/support"
	"sync/atomic"
)

func updateStats(direction string, totalFilled int,  fills []support.Fill, cash *int64, position *uint64) {
	if direction == "buy" {
		pos := atomic.AddUint64(position, uint64(totalFilled))
		fmt.Printf("position: %v\n", pos)

		value, _ := totalFills(fills)
		n := atomic.AddInt64(cash, int64(value*-1))
		fmt.Printf("cash: %v\n", n)
	} else {
		pos := atomic.AddUint64(position, ^uint64(totalFilled-1))
		fmt.Printf("position: %v\n", pos)

		value, _ := totalFills(fills)
		n := atomic.AddInt64(cash, int64(value))
		fmt.Printf("cash: %v\n", n)
	}
}

func orderStatus(openOrders chan support.Order, cash *int64, position *uint64) {
	for {
		select {
		case o := <-openOrders:
			status, err := support.GetStatus(o.Id)
			if err == nil {
				if !status.Open {
					fmt.Printf("order id:%v closed\n", o.Id)
					fmt.Println(status.Fills)

					updateStats(o.Direction, status.TotalFilled,  status.Fills, cash, position)

				} else {
					//fmt.Printf("order not closed; putting order back on channel")
					openOrders <- o
				}
			}
		default:
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func totalFills(fills []support.Fill) (total int, totalFilled int){
	for _, f := range fills {
		total += f.Price*f.Qty
		totalFilled += f.Qty
	}

	return total, totalFilled
}

func tallyExistingOrders(allOrders support.AllOrders, cash *int64, position *uint64) {
	for _, o := range allOrders.Orders {
		if !o.Open {
			updateStats(o.Direction, o.TotalFilled,  o.Fills, cash, position)
		}
	}
}

func main() {
	sconf.Must(&support.Cfg).Read(ini.File("config.gcfg"))
	fmt.Println(support.Cfg.Stockfighter.ApiKey)
	fmt.Println(support.Cfg.Stockfighter.Account)
	fmt.Println(support.Cfg.Stockfighter.Venue)
	fmt.Println(support.Cfg.Stockfighter.Symbol)
	fmt.Println(support.Cfg.Stockfighter.BaseUrl)

	//lowPrice := flag.Int("low", 0, "price to prime engine with")
	//flag.Parse()

	var cash int64
	var position uint64

	//smallBlock := 5

	//openOrders := make(chan support.Order, 100)

	allOrders, err := support.GetStatusForStock()
	if err == nil {
		tallyExistingOrders(allOrders, &cash, &position)
	}


	//go orderStatus(openOrders, &cash, &position)
	//
	//for i:=0; i<1; i++ {
	//	support.Buy("limit", *lowPrice, smallBlock, openOrders)
	//	time.Sleep(3 * time.Second)
	//}
}
