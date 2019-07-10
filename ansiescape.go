package ansiescape

import "fmt"

const esc = "\x1b"

// ClearDisplay returns the key to clear the display
// n=0 clear after the cursor
// n=1 clear before the cursor
// n=2 clear all
func ClearDisplay(n int) string {
	return fmt.Sprintf("%v[%dJ", esc, n)
}

// ClearRow returns the key to clear the row
// n=0 clear the row of after the cursor
// n=1 clear the row of before the cursor
// n=2 clear row all
func ClearRow(n int) string {
	return fmt.Sprintf("%v[%dK", esc, n)
}
