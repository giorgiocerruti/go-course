package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

type menuItem struct {
	name   string
	prices map[string]float64
}

type Menu []menuItem

func (m Menu) print() {
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))

		for size, cost := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, cost)
		}
	}
}

func (m *Menu) add() {
	fmt.Print("Add a new item: ")

	name, err := in.ReadString('\n')
	if err != nil {
		panic("Error reading streing")
	}

	*m = append(*m, menuItem{name: strings.TrimSpace(name), prices: make(map[string]float64)})
}
func Add() {
	data.add()
}

func Print() {
	data.print()
}
