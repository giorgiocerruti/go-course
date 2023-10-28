package main

import "fmt"

const (
	statusNew = iota
	statusValidated
	statusProcessed
	statusFailed
	unknownStatus
)

var stringStatus = map[orderStatus]string{
	statusNew:       "New",
	statusValidated: "Validated",
	statusProcessed: "Processed",
	statusFailed:    "Failed",
	unknownStatus:   "Unknown",
}

type order struct {
	ProductCode int
	Quantity    float64
	Status      orderStatus
}

type orderStatus int

func (o order) getStatus() string {
	return stringStatus[o.Status]
}

func (o order) String() string {
	return fmt.Sprintf("ProductCode: %d, Quantity: %f, Status: %s\n",
		o.ProductCode, o.Quantity, o.getStatus())
}

type orderValidation struct {
	order order
	err   error
}

var orders []order
