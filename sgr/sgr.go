package sgr

import (
	"fmt"
	"strings"
)

// SGR is Select Graphic Rendition command type.
// for graphic ansi escape code.
type SGR string

const (
	// Clear is all specified SGR code.
	Clear SGR = "0"
	// Bold is for making the text to bold
	Bold SGR = "1"
	// Light is for making the text to light color
	Light SGR = "2"
	// Italic is for making the text to itelic
	Italic SGR = "3"
	// Underline is for drwing the underline
	Underline SGR = "4"
	// Blink is for blinking the text
	Blink SGR = "5"
	// FastBlink is for blinking the text
	FastBlink SGR = "6"
	// Reverse is for reversing the text foreground to background
	Reverse SGR = "7"
	// Hide is for hiding
	Hide SGR = "8"
	// Undo is for undoing
	Undo SGR = "9"
	// FgBlack is ANSI color Black
	FgBlack SGR = "30"
	// FgRed is ANSI color Red
	FgRed SGR = "31"
	// FgGreen is ANSI color Green
	FgGreen SGR = "32"
	// FgYellow is ANSI color Yellow
	FgYellow SGR = "33"
	// FgBlue is ANSI color Blue
	FgBlue SGR = "34"
	// FgMagenta is ANSI color Magenta
	FgMagenta SGR = "35"
	// FgCyan is ANSI color Cyan
	FgCyan SGR = "36"
	// FgWhite is ANSI color White
	FgWhite SGR = "37"
	// FgExColor is ANSI color for extentions
	fgExColor SGR = "38;5;"
	// ClearFgColor is to back the default color
	ClearFgColor SGR = "39"
	// BgBlack is ANSI color Black
	BgBlack SGR = "40"
	// BgRed is ANSI color Red
	BgRed SGR = "41"
	// BgGreen is ANSI color Green
	BgGreen SGR = "42"
	// BgYellow is ANSI color Yellow
	BgYellow SGR = "43"
	// BgBlue is ANSI color Blue
	BgBlue SGR = "44"
	// BgMagenta is ANSI color Magenta
	BgMagenta SGR = "45"
	// BgCyan is ANSI color Cyan
	BgCyan SGR = "46"
	// BgWhite is ANSI color White
	BgWhite SGR = "47"
	// BgExColor is ANSI color for extentions
	bgExColor SGR = "48;5;"
	// ClearBgColor is to back the default color
	ClearBgColor SGR = "49"
	// HFgBlack is ANSI color Black(hign contrast)
	HFgBlack SGR = "90"
	// HFgRed is ANSI color Red(hign contrast)
	HFgRed SGR = "91"
	// HFgGreen is ANSI color Green(hign contrast)
	HFgGreen SGR = "92"
	// HFgYellow is ANSI color Yellow(hign contrast)
	HFgYellow SGR = "93"
	// HFgBlue is ANSI color Blue(hign contrast)
	HfgColor4 SGR = "94"
	// HFgMagenta is ANSI color Magenta(hign contrast)
	HFgMagenta SGR = "95"
	// HFgCyan is ANSI color Cyan(hign contrast)
	HFgCyan SGR = "96"
	// HFgWhite is ANSI color White(hign contrast)
	HFgWhite SGR = "97"
	// HBgBlack is ANSI color Black(hign contrast)
	HBgBlack SGR = "100"
	// HBgRed is ANSI color Red(hign contrast)
	HBgRed SGR = "101"
	// HBgGreen is ANSI color Green(hign contrast)
	HBgGreen SGR = "102"
	// HBgYellow is ANSI color Yellow(hign contrast)
	HBgYellow SGR = "103"
	// HBgBlue is ANSI color Blue(hign contrast)
	HBgBlue SGR = "104"
	// HBgMagenta is ANSI color Magenta(hign contrast)
	HBgMagenta SGR = "105"
	// HBgCyan is ANSI color Cyan(hign contrast)
	HBgCyan SGR = "106"
	// HBgWhite is ANSI color White(hign contrast)
	HBgWhite SGR = "107"
)

const (
	esc = "\x1b"
)

// FgColor returns 256 ANSI color for foreground color(text color)
// x is color index(0 ~ 256)
func FgColor(x int) SGR {
	return SGR(fmt.Sprintf("%v%v", fgExColor, x))
}

// BgColor returns 256 ANSI color for background
// x is color index(0 ~ 256)
func BgColor(x int) SGR {
	return SGR(fmt.Sprintf("%v%v", bgExColor, x))
}

// EscSeq returns created ANSI Escape sequence by sgrs
func EscSeq(sgrs ...SGR) string {
	return fmt.Sprintf("%v[%vm", esc, join(sgrs))
}

func join(sgrs []SGR) string {
	return strings.Join(convertToString(sgrs), ";")
}

func convertToString(sgrs []SGR) []string {
	strs := make([]string, 0, len(sgrs))
	for _, sgr := range sgrs {
		strs = append(strs, string(sgr))
	}
	return strs
}
