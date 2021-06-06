package withstyle

// Red color of text is rgb(229, 57, 53).
func Red(s string) string { return Rgb(229, 57, 53)(s) }

// Pink color of text is rgb(216, 27, 96)).
func Pink(s string) string { return Rgb(216, 27, 96)(s) }

// Purple color of text is rgb(142, 36, 170).
func Purple(s string) string { return Rgb(142, 36, 170)(s) }

// DeepPurple color of text is rgb(94, 53, 177).
func DeepPurple(s string) string { return Rgb(94, 53, 177)(s) }

// Indigo color of text is rgb(57, 73, 171).
func Indigo(s string) string { return Rgb(57, 73, 171)(s) }

// GrayBlue color of text is rgb(30, 136, 229).
func GrayBlue(s string) string { return Rgb(30, 136, 229)(s) }

// LightBlue color of text is rgb(3, 155, 229).
func LightBlue(s string) string { return Rgb(3, 155, 229)(s) }

// Cyan color of text is rgb(0, 151, 167).
func Cyan(s string) string { return Rgb(0, 151, 167)(s) }

// Teal color of text is rgb(0, 137, 123).
func Teal(s string) string { return Rgb(0, 137, 123)(s) }

// Green color of text is rgb(67, 160, 71).
func Green(s string) string { return Rgb(67, 160, 71)(s) }

// LightGreen color of text is rgb(124, 179, 66).
func LightGreen(s string) string { return Rgb(124, 179, 66)(s) }

// Lime color of text is rgb(192, 202, 51).
func Lime(s string) string { return Rgb(192, 202, 51)(s) }

// Yellow color of text is rgb(252, 215, 52).
func Yellow(s string) string { return Rgb(252, 215, 52)(s) }

// Amber color of text is rgb(255, 179, 0).
func Amber(s string) string { return Rgb(255, 179, 0)(s) }

// Orange color of text is rgb(251, 140, 0).
func Orange(s string) string { return Rgb(251, 140, 0)(s) }

// DeepOrange color of text is rgb(244, 81, 30).
func DeepOrange(s string) string { return Rgb(244, 81, 30)(s) }

// Black color of text is rgb(0, 0, 0).
func Black(s string) string { return Rgb(0, 0, 0)(s) }

// DarkGray color of text is rgb(117, 117, 117).
func DarkGray(s string) string { return Rgb(117, 117, 117)(s) }

// Gray color of text is rgb(189, 189, 189).
func Gray(s string) string { return Rgb(189, 189, 189)(s) }

// White color of text is rgb(245, 245, 245).
func White(s string) string { return Rgb(245, 245, 245)(s) }

// BgRed background color is rgb(229, 57, 53).
func BgRed(s string) string { return BgRgb(229, 57, 53)(s) }

// BgPink background color is rgb(216, 27, 96).
func BgPink(s string) string { return BgRgb(216, 27, 96)(s) }

// BgPurple background color is rgb(142, 36, 170).
func BgPurple(s string) string { return BgRgb(142, 36, 170)(s) }

// BgDeepPurple background color is rgb(94, 53, 177).
func BgDeepPurple(s string) string { return BgRgb(94, 53, 177)(s) }

// BgIndigo background color is rgb(57, 73, 171).
func BgIndigo(s string) string { return BgRgb(57, 73, 171)(s) }

// BgBlue background color is rgb(30, 136, 229).
func BgBlue(s string) string { return BgRgb(30, 136, 229)(s) }

// BgLightBlue background color is rgb(3, 155, 229).
func BgLightBlue(s string) string { return BgRgb(3, 155, 229)(s) }

// BgCyan background color is rgb(0, 151, 167).
func BgCyan(s string) string { return BgRgb(0, 151, 167)(s) }

// BgTeal background color is rgb(0, 137, 123).
func BgTeal(s string) string { return BgRgb(0, 137, 123)(s) }

// BgGreen background color is rgb(67, 160, 71).
func BgGreen(s string) string { return BgRgb(67, 160, 71)(s) }

// BgLightGreen background color is rgb(124, 179, 66).
func BgLightGreen(s string) string { return BgRgb(124, 179, 66)(s) }

// BgLime background color is rgb(192, 202, 51).
func BgLime(s string) string { return BgRgb(192, 202, 51)(s) }

// BgYellow background color is rgb(252, 215, 52).
func BgYellow(s string) string { return BgRgb(252, 215, 52)(s) }

// BgAmber background color is rgb(255, 179, 0).
func BgAmber(s string) string { return BgRgb(255, 179, 0)(s) }

// BgOrange background color is rgb(251, 140, 0).
func BgOrange(s string) string { return BgRgb(251, 140, 0)(s) }

// BgDeepOrange background color is rgb(244, 81, 30).
func BgDeepOrange(s string) string { return BgRgb(244, 81, 30)(s) }

// BgBlack background color is rgb(0, 0, 0).
func BgBlack(s string) string { return BgRgb(0, 0, 0)(s) }

// BgDarkGray background color is rgb(117, 117, 117).
func BgDarkGray(s string) string { return BgRgb(117, 117, 117)(s) }

// BgGray background color is rgb(245, 245, 245).
func BgGray(s string) string { return BgRgb(189, 189, 189)(s) }

// BgWhite background color is rgb(245, 245, 245).
func BgWhite(s string) string { return BgRgb(245, 245, 245)(s) }

// Bold font-style makes text little wider.
func Bold(s string) string { return "\033[01m" + s + reset }

// Dim decreases intensity of text.
func Dim(s string) string { return "\033[02m" + s + reset }

// Italic font-style makes text slant slightly to the right.
func Italic(s string) string { return "\033[03m" + s + reset }

// Underline text-decoration adds a stroke under text.
func Underline(s string) string { return "\033[04m" + s + reset }

// Reverse text-decoration swaps text and bg colors.
func Reverse(s string) string { return "\033[07m" + s + reset }

// Hidden text-decoration hides text.
func Hidden(s string) string { return "\033[08m" + s + reset }
