package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var receivedOrderCh = make(chan order)
	var validOrderCh = make(chan order)
	var invalidOrdersCh = make(chan orderValidation)

	go receiveOrders(receivedOrderCh)
	go validateOrders(receivedOrderCh, validOrderCh, invalidOrdersCh)
	wg.Add(1)

	go func(validOrderCh <-chan order, invalidOrdersCh <-chan orderValidation) {
	loop:
		for {
			select {

			case validOrder, ok := <-validOrderCh:
				if !ok {
					break loop
				}
				fmt.Println("Valid order:", validOrder)
			case order, ok := <-invalidOrdersCh:
				if !ok {
					break loop
				}
				fmt.Printf("Error: %v - Order %v", order.err, order.order)
			}
		}

		wg.Done()
	}(validOrderCh, invalidOrdersCh)

	wg.Wait()

}

func receiveOrders(in chan<- order) {

	for _, rawOrder := range newOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			fmt.Println(err)
			continue
		}
		in <- newOrder
	}

	close(in)
}

func validateOrders(in <-chan order, out chan<- order, errCh chan orderValidation) {

	for order := range in {
		if order.Quantity <= 0 {
			errCh <- orderValidation{order: order, err: fmt.Errorf("Invalid quantity")}
		} else {
			out <- order
		}
	}

	defer close(out)
	defer close(errCh)
}

var newOrders = []string{
	`{"productCode": 123, "quantity": -1, "status": 0}`,
	`{"productCode": 134, "quantity": 3, "status": 0}`,
	`{"productCode": 222, "quantity": 5, "status": 0}`,
	`{"productCode": 333, "quantity": 5, "status": 0}`,
	`{"productCode": 444, "quantity": 5, "status": 0}`,
}
