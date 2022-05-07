package utils

import "gotham/helpers"

type IOrder interface {
	Get() *Order
	GetOrderBy() string
	GetSortBy() string
}

type Order struct {
	OrderBy string `query:"order_by"`
	SortBy  string `query:"sort_by"`
}

func (o *Order) Get() *Order {
	return o
}

func (o *Order) GetOrderBy() string {
	return o.OrderBy
}

func (o *Order) GetSortBy() string {
	if !helpers.InArray(o.SortBy, []string{"asc", "desc"}) {
		o.SortBy = "asc"
	}
	return o.SortBy
}
