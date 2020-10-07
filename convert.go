package withstyle

import (
	"strconv"
)

// Be careful! A lot of Magic Numbers inside
func rgbToAnsi256(r, g, b int) string {
	var ansi int
	if r == g && g == b {
		if r < 8 {
			ansi = 16
		} else if r > 248 {
			ansi = 231
		} else {
			ansi = 232 + 24*intRound((float64(r)-8)/247)
		}
	}

	ansi = 16 + (36 * intRound(float64(r)/255*5)) + (6 * intRound(float64(g)/255*5)) + intRound(float64(b)/255*5)

	return strconv.Itoa(ansi)
}
