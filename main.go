package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/giorgiocerruti/demo/menu"
)

func main() {
	var in = bufio.NewReader(os.Stdin)

loop:
	for {
		fmt.Println("Please select your option")
		fmt.Println("1) Print Menu")
		fmt.Println("2) Add a new item")
		fmt.Println("q) quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.Print()
		case "2":
			menu.Add()
		case "q":
			fmt.Println("Bye!")
			break loop
		default:
			fmt.Println("Uknow option")
		}
	}

}
