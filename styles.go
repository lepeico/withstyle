package withstyle

func Red(s string) string {
	return Rgb(229, 57, 53)(s)
}

func Pink(s string) string {
	return Rgb(216, 27, 96)(s)
}

func Purple(s string) string {
	return Rgb(142, 36, 170)(s)
}

func DeepPurple(s string) string {
	return Rgb(94, 53, 177)(s)
}

func Indigo(s string) string {
	return Rgb(57, 73, 171)(s)
}

func Blue(s string) string {
	return Rgb(30, 136, 229)(s)
}

func LightBlue(s string) string {
	return Rgb(3, 155, 229)(s)
}

func Cyan(s string) string {
	return Rgb(0, 151, 167)(s)
}

func Teal(s string) string {
	return Rgb(0, 137, 123)(s)
}

func Green(s string) string {
	return Rgb(67, 160, 71)(s)
}

func LightGreen(s string) string {
	return Rgb(124, 179, 66)(s)
}

func Lime(s string) string {
	return Rgb(192, 202, 51)(s)
}

func Yellow(s string) string {
	return Rgb(252, 215, 52)(s)
}

func Amber(s string) string {
	return Rgb(255, 179, 0)(s)
}

func Orange(s string) string {
	return Rgb(251, 140, 0)(s)
}

func DeepOrange(s string) string {
	return Rgb(244, 81, 30)(s)
}

func Black(s string) string {
	return Rgb(0, 0, 0)(s)
}

func DarkGray(s string) string {
	return Rgb(117, 117, 117)(s)
}

func Gray(s string) string {
	return Rgb(189, 189, 189)(s)
}

func White(s string) string {
	return Rgb(245, 245, 245)(s)
}

func BgRed(s string) string {
	return BgRgb(229, 57, 53)(s)
}

func BgPink(s string) string {
	return BgRgb(216, 27, 96)(s)
}

func BgPurple(s string) string {
	return BgRgb(142, 36, 170)(s)
}

func BgDeepPurple(s string) string {
	return BgRgb(94, 53, 177)(s)
}

func BgIndigo(s string) string {
	return BgRgb(57, 73, 171)(s)
}

func BgBlue(s string) string {
	return BgRgb(30, 136, 229)(s)
}

func BgLightBlue(s string) string {
	return BgRgb(3, 155, 229)(s)
}

func BgCyan(s string) string {
	return BgRgb(0, 151, 167)(s)
}

func BgTeal(s string) string {
	return BgRgb(0, 137, 123)(s)
}

func BgGreen(s string) string {
	return BgRgb(67, 160, 71)(s)
}

func BgLightGreen(s string) string {
	return BgRgb(124, 179, 66)(s)
}

func BgLime(s string) string {
	return BgRgb(192, 202, 51)(s)
}

func BgYellow(s string) string {
	return BgRgb(252, 215, 52)(s)
}

func BgAmber(s string) string {
	return BgRgb(255, 179, 0)(s)
}

func BgOrange(s string) string {
	return BgRgb(251, 140, 0)(s)
}

func BgDeepOrange(s string) string {
	return BgRgb(244, 81, 30)(s)
}

func BgBlack(s string) string {
	return BgRgb(0, 0, 0)(s)
}

func BgDarkGray(s string) string {
	return BgRgb(117, 117, 117)(s)
}

func BgGray(s string) string {
	return BgRgb(189, 189, 189)(s)
}

func BgWhite(s string) string {
	return BgRgb(245, 245, 245)(s)
}

func Bold(s string) string {
	return "\033[01m" + s + reset
}

func Dim(s string) string {
	return "\033[02m" + s + reset
}

func Italic(s string) string {
	return "\033[03m" + s + reset
}

func Underline(s string) string {
	return "\033[04m" + s + reset
}

func Reverse(s string) string {
	return "\033[07m" + s + reset
}

func Hidden(s string) string {
	return "\033[08m" + s + reset
}
