package support

import (
	"fmt"
	"github.com/franela/goreq"
)

func Buy(orderType string, price int64, qty int, openOrders chan Order) {
	order(orderType, "buy", price, qty, openOrders)
}

func Sell(orderType string, price int64, qty int, openOrders chan Order) {
	order(orderType, "sell", price, qty, openOrders)
}

func order(orderType string, direction string, price int64, qty int, openOrders chan Order) {
	order := Order{
		Account:   Cfg.Stockfighter.Account,
		Venue:     Cfg.Stockfighter.Venue,
		Symbol:    Cfg.Stockfighter.Symbol,
		Price:     price,
		Qty:       qty,
		Direction: direction,
		OrderType: orderType,
	}

	uri := fmt.Sprintf("%s/venues/%s/stocks/%s/orders", Cfg.Stockfighter.BaseUrl, order.Venue, order.Symbol)
	req := goreq.Request{
		Method: "POST",
		Uri:    uri,
		Body:   order,
	}

	req.AddHeader("X-Starfighter-Authorization", Cfg.Stockfighter.ApiKey)
	res, err := req.Do()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		if res.StatusCode == 200 {
			res.Body.FromJsonTo(&order)
			fmt.Println(order)
			openOrders <- order
		}
	}
}
