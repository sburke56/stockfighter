package main

import (
	"fmt"
	"github.com/franela/goreq"
)

func buy(orderType string, price int, qty int) {
	order := Order{
		Account:   Cfg.Stockfighter.Account,
		Venue:     Cfg.Stockfighter.Venue,
		Symbol:    Cfg.Stockfighter.Symbol,
		Price:     price,
		Qty:       qty,
		Direction: "buy",
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
		fmt.Println(res.Body.ToString())
	}
}
