package withstyle

import (
	"strconv"
)

// Be careful! A lot of Magic Numbers inside
func rgbToAnsi256(r, g, b float64) string {
	var ansi int
	if r == g && g == b {
		if r < 8 {
			ansi = 16
		} else if r > 248 {
			ansi = 231
		} else {
			ansi = 232 + 24*intRound((r-8)/247)
		}
	}

	ansi = 16 + (36 * intRound(r/255*5)) + (6 * intRound(g/255*5)) + intRound(b/255*5)

	return strconv.Itoa(ansi)
}

func rgbToAnsi16(r, g, b float64) string {
	var ansi int

	shade := intRound(max(r/255, g/255, b/255) * 2)

	if shade == 0 {
		ansi = 0
	} else {
		ansi = ((intRound(b/255) << 2) | (intRound(g/255) << 1) | intRound(r/255))
	}

	if shade == 2 {
		ansi += 8
	}

	return strconv.Itoa(ansi)
}
