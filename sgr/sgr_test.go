package sgr

import "testing"

func TestEscSeq(t *testing.T) {
	tests := []struct {
		exp string
		act string
	}{
		{"\x1b[1m", MakeString(Bold)},
		{"\x1b[1;30;47m", MakeString(Bold, FgBlack, BgWhite)},
		{"\x1b[38;5;120m", MakeString(FgColor(120))},
		{"\x1b[38;5;120;48;5;5m", MakeString(FgColor(120), BgColor(5))},
		{"\x1b[38;5;120;3;48;5;5m", MakeString(FgColor(120), Italic, BgColor(5))},
	}

	for _, test := range tests {
		if test.exp != test.act {
			t.Fatalf("expected is %v but actual is %v", test.exp, test.act)
		}
	}
}
