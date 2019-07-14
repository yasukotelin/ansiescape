package ansiescape

import "fmt"

const esc = "\x1b"

func makeEscSequence(n int, op string) string {
	return fmt.Sprintf("\x1b[%d%v", n, op)
}

// MoveCursorUp returns the escape sequence string to move
// the cursor up n.
func MoveCursorUp(n int) string {
	return makeEscSequence(n, "A")
}

// MoveCursorDown returns the escape sequence string to move
// the cursor down n.
func MoveCursorDown(n int) string {
	return makeEscSequence(n, "B")
}

// MoveCursorRight returns the escape sequence string to move
// the cussor right n.
func MoveCursorRight(n int) string {
	return makeEscSequence(n, "C")
}

// MoveCursorLeft returns the escape sequence string to move
// the cussor left n.
func MoveCursorLeft(n int) string {
	return makeEscSequence(n, "D")
}

// MoveFromLeftEnd returns the escape sequence string to move
// the cursor n from left end
func MoveFromLeftEnd(n int) string {
	return makeEscSequence(n, "G")
}

// MoveAbsolute returns the escape sequence string to move
// the cursor absolute
func MoveAbsolute(x, y int) string {
	return fmt.Sprintf("\x1b[%d;%dH", x, y)
}

// ClearDisplay returns the escape sequence string to clear the display
// n=0 clear after the cursor
// n=1 clear before the cursor
// n=2 clear all
func ClearDisplay(n int) string {
	return fmt.Sprintf("%v[%dJ", esc, n)
}

// ClearRow returns the escape sequence string to clear the row
// n=0 clear the row of after the cursor
// n=1 clear the row of before the cursor
// n=2 clear row all
func ClearRow(n int) string {
	return fmt.Sprintf("%v[%dK", esc, n)
}

// ScrollNext returns the escape sequence string to scroll next
func ScrollNext(n int) string {
	return makeEscSequence(n, "S")
}

// ScrollPrev returns the escape sequence string to scroll previous
func ScrollPrev(n int) string {
	return makeEscSequence(n, "T")
}
