package main

import (
	"fmt"

	"github.com/yasukotelin/ansiescape/sgr"
)

func main() {
	clear := sgr.EscSeq(sgr.Clear)
	seq := sgr.EscSeq(sgr.BgColor(55), sgr.FgWhite)

	fmt.Printf("%vhello, world%v", seq, clear)
}
