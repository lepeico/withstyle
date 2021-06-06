package withstyle

import (
	"math"
	"strconv"
)

const (
	colorStart = "\033[38;05;"
	bgStart    = "\033[48;05;"
	reset      = "\033[0m"
)

type StyleFn func(string) string

func WithStyles(styles ...StyleFn) StyleFn {
	return func(hex string) string {
		style := styles[0]

		if len(styles) == 1 {
			return style(hex)
		}

		return style(WithStyles(styles[1:]...)(hex))
	}
}

func Rgb(r, g, b float64) StyleFn {
	return func(str string) string {
		return colorStart + rgbToAnsi256(r, g, b) + "m" + str + reset
	}
}

func BgRgb(r, g, b float64) StyleFn {
	return func(str string) string {
		return bgStart + rgbToAnsi256(r, g, b) + "m" + str + reset
	}
}

func Hex(hex string) StyleFn {
	r, g, b := hex2rgb(hex)

	return func(str string) string {
		return colorStart + rgbToAnsi256(r, g, b) + "m" + str + reset
	}
}

func BgHex(hex string) StyleFn {
	r, g, b := hex2rgb(hex)

	return func(str string) string {
		return bgStart + rgbToAnsi256(r, g, b) + "m" + str + reset
	}
}

func hex2rgb(hex string) (r, g, b float64) {
	if hex[0] != '#' {
		return
	}

	hex2byte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}

		return 0
	}

	switch len(hex) {
	case 7:
		r = float64(hex2byte(hex[1])<<4 + hex2byte(hex[2]))
		g = float64(hex2byte(hex[3])<<4 + hex2byte(hex[4]))
		b = float64(hex2byte(hex[5])<<4 + hex2byte(hex[6]))
	case 4:
		r = float64(hex2byte(hex[1]) * 17)
		g = float64(hex2byte(hex[2]) * 17)
		b = float64(hex2byte(hex[3]) * 17)
	default:
		return
	}

	return r, g, b
}

func rgbToAnsi256(r, g, b float64) string {
	ansi := 16 +
		(36 * math.Round(r/255*5)) +
		(6 * math.Round(g/255*5)) +
		math.Round(b/255*5)

	return strconv.Itoa(int(ansi))
}
