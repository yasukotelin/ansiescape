package main

import (
	"fmt"

	"github.com/yasukotelin/ansiescape"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
		if i == 4 {
			fmt.Print(ansiescape.ClearRow(2))
		}
	}
}
