package support

var Cfg Config

type Config struct {
	Stockfighter struct {
		ApiKey  string
		Account string
		Symbol  string
		Venue   string
		BaseUrl string
	}
}

type Fill struct {
	Price int
	Qty int
}

type Order struct {
	Id int
	Account   string
	Venue     string
	Symbol    string
	Price     int64
	Qty       int
	Direction string
	OrderType string
	Fills []Fill
	Open bool
	TotalFilled int64
}

type Status struct {
	Direction string
	OriginalQty int
	Qty int
	Price int
	OrderType string
	Id int
	Fills []Fill
	TotalFilled int
	Open bool
}

type AllOrders struct {
	Orders []Order
}

type Quote struct {
	Ask int
	Last int64
}

type BidAsk struct {
	Price int
	Qty   int
	IsBuy bool
}

type OrderBook struct {
	Venue  string
	Symbol string
	Bids   []BidAsk
	Asks   []BidAsk
}

type Position struct {
	Remaining   int
	Outstanding int
	Filled      int
}
