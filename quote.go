package main

import (
	"fmt"
	"github.com/franela/goreq"
)

func getQuote(venue string, stock string) (quote Quote, err error) {
	uri := fmt.Sprintf("%s/venues/%s/stocks/%s/quote", Cfg.Stockfighter.BaseUrl, Cfg.Stockfighter.Venue, Cfg.Stockfighter.Symbol)
	req := goreq.Request{Uri: uri}
	req.AddHeader("X-Starfighter-Authorization", Cfg.Stockfighter.ApiKey)

	res, err := req.Do()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if res.StatusCode == 200 {
			res.Body.FromJsonTo(&quote)
		}
	}

	return quote, err
}
