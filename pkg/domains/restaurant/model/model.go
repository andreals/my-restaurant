package model

import (
	"time"
)

type Restaurant struct {
	Name  string
	Chefs int64
	Menu  Menu
}

type Menu struct {
	Dishes []Dish
}

type Table struct {
	Number    int64
	Customers int64
	Orders    []Order
}

type Order struct {
	Customer int64
	Dishes   []Dish
}

type Dish struct {
	Name            string
	Price           float64
	PreparationTime time.Duration
}
