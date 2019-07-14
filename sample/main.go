package main

import (
	"fmt"
	"time"

	"github.com/yasukotelin/ansiescape"
)

func main() {
	fmt.Print("hello, world")

	wait1Seconds()
	clearRowAndReset()

	fmt.Println("clear row")
}

func wait1Seconds() {
	time.Sleep(1 * time.Second)
}

func clearRowAndReset() {
	fmt.Print(ansiescape.ClearRow(2), ansiescape.MoveFromLeftEnd(0))
}
