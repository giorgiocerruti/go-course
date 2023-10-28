package main

import (
	"bytes"
	"fmt"
	"strings"
)

type printer interface {
	Print() string
}

type person struct {
	name string
	age  int32
}

func (p person) Print() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}

type maneuItem struct {
	name   string
	prices map[string]float64
}

func (m maneuItem) Print() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Name: %s\n", m.name))
	b.WriteString(strings.Repeat("-", 10))
	for size, cost := range m.prices {
		b.WriteString(fmt.Sprintf("%s: %f\n", size, cost))
	}

	return b.String()
}

func Print(p printer) {
	fmt.Println("Type:", )
	fmt.Println(p.Print())
}

func main() {
	var p printer

	p = person{name: "Giorgio", age: 30}
	fmt.Println(p.Print())

	p = maneuItem{
		name: "Pizza",
		prices: map[string]float64{
			"small":  5.99,
			"medium": 7.99,
			"large":  9.99,
		},
	}

	fmt.Println(p.Print())

	switch c := p.(type) {
	case person:
		fmt.Println("Person", c.name)
	case maneuItem:
		fmt.Println("ManeuItem", c.name)
	default:
		fmt.Println("Unknown type")
	}

	fmt.Println("Using Print function")

	Print(person{name: "Giorgio", age: 30})
	Print(maneuItem{name: "Pizza", prices: map[string]float64{"small": 5.99, "medium": 7.99, "large": 9.99}})
}
