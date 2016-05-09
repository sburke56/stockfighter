package support

import (
	"fmt"
	"github.com/franela/goreq"
)

func GetStatus(id int) (status Status, err error){
	uri := fmt.Sprintf("%s/venues/%s/stocks/%s/orders/%d",
		Cfg.Stockfighter.BaseUrl,
		Cfg.Stockfighter.Venue,
		Cfg.Stockfighter.Symbol,
		id)
	req := goreq.Request{ Uri: uri }

	req.AddHeader("X-Starfighter-Authorization", Cfg.Stockfighter.ApiKey)

	res, err := req.Do()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if res.StatusCode == 200 {
			res.Body.FromJsonTo(&status)
		}
	}

	return status, err
}

func GetStatusForStock() (allOrders AllOrders, err error){
	uri := fmt.Sprintf("%s/venues/%s/accounts/%s/stocks/%s/orders",
		Cfg.Stockfighter.BaseUrl,
		Cfg.Stockfighter.Venue,
		Cfg.Stockfighter.Account,
		Cfg.Stockfighter.Symbol)
	req := goreq.Request{ Uri: uri }

	req.AddHeader("X-Starfighter-Authorization", Cfg.Stockfighter.ApiKey)

	res, err := req.Do()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if res.StatusCode == 200 {
			res.Body.FromJsonTo(&allOrders)
		}
	}

	return allOrders, err
}
