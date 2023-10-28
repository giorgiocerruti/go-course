package main

import (
	_ "embed"
	"fmt"
	"strings"
	"strconv"
	"github.com/giorgiocerruti/demo1/calc"
	"github.com/pioz/faker"
)

var (
	//go:embed numbers.txt
	data []byte
)

func init() {
	fmt.Println("Hello Josh!")
}


func main() {
	var price, disc int

	price = 4
	disc = 2

	total := calc.Discount(price, disc)

	fmt.Println("Price: ", price)
	fmt.Println("Discount: ", total)

	fmt.Println("Total: ", total)
	fmt.Println("Random Color: ", faker.ColorName())


	fmt.Println(strings.Repeat("-", 30))

	fmt.Println(string(data))

	var sum int

	for _, line := range strings.Split(string(data), "\n\r") {
		if line != "" {
			val,_ := strconv.Atoi(line)
			fmt.Println(val)
			sum += val
		}
	}

	fmt.Println("Sum of all numbers:", sum)

}
