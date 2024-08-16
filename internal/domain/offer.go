package domain

import "time"

type League = string
type Quantity = int

type Base struct {
	Quantity Quantity
	Symbol   Symbol
}

type Quote struct {
	Quantity Quantity
	Stock    Quantity
	Symbol   Symbol
}

type Offer struct {
	Base      Base
	Quote     Quote
	League    League
	Timestamp time.Time
}

type OffersRepository interface {
	Add(offer Offer) error
}
