package main

import (
	"fmt"

	"github.com/yasukotelin/ansiescape/sgr"
)

func main() {
	clear := sgr.MakeString(sgr.Clear)
	seq := sgr.MakeString(sgr.BgColor(55), sgr.FgWhite)

	fmt.Printf("%vhello, world%v", seq, clear)
}
