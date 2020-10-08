package withstyle

const (
	reset = "\033[0m"
)

type styleFn func(string) string

func WithStyles(styles ...styleFn) styleFn {
	return func(s string) string {
		style := styles[0]

		if len(styles) == 1 {
			return style(s)
		}

		return style(WithStyles(styles[1:len(styles)]...)(s))
	}
}

func Rgb(r, g, b float64) func(string) string {
	return func(str string) string {
		return "\033[38;05;" + rgbToAnsi256(r, g, b) + "m" + str + reset
	}
}

func BgRgb(r, g, b float64) func(string) string {
	return func(str string) string {
		return "\033[48;05;" + rgbToAnsi256(r, g, b) + "m" + str + reset
	}
}
