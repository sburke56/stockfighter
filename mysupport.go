package main

type Order struct {
	Account string
	Venue string
	Symbol string
	Price int
	Qty int
	Direction string
	OrderType string
}

type Quote struct {
	Ask int
}

type BidAsk struct {
	Price int
	Qty int
	IsBuy bool
}

type OrderBook struct {
	Venue string
	Symbol string
	Bids [] BidAsk
	Asks [] BidAsk
}

type Position struct {
	Remaining int
	Outstanding int
	Filled int
}
