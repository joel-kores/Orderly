package models

type Order struct {
	CommonFields
	CustomerID uint
	Customer   Customer
	TotalPrice float64
}
