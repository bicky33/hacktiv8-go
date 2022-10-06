package domain

import (
	"time"
)

type Order struct {
	OrderId      int
	CustomerName string
	OrderedAt    time.Time
	Items        []Item
}

type Item struct {
	ItemId      int
	ItemCode    string
	Description string
	Quantity    int
	OrderId     int
}
