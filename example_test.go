package withstyle_test

import (
	"fmt"

	w "github.com/lepeico/withstyle"
)

func ExampleWithStyles() {
	// Define theme
	var (
		debug   = w.WithStyles(w.BgWhite, w.Black, w.Italic)
		err     = w.WithStyles(w.BgRed, w.Black, w.Bold)
		warning = w.WithStyles(w.BgOrange, w.White, w.Underline)
		success = w.WithStyles(w.BgGreen, w.White, w.Italic)
		info    = w.WithStyles(w.BgBlue, w.Black, w.Italic)
		verbose = w.WithStyles(w.BgDeepPurple, w.White)
		fatal   = w.WithStyles(w.BgBlack, w.Red, w.Bold)
	)

	fmt.Println(debug("Okay. This is a debug message!"))
	fmt.Println(err("Oops! This is an error message!"))
	fmt.Println(warning("Hm.. This is a warning message!"))
	fmt.Println(success("Yay! This is a success message!"))
	fmt.Println(info("Hi! Just an info message!"))
	fmt.Println(verbose("So.. This is a verbose message!"))
	fmt.Println(fatal("%*#$! This is a fatal message!"))
}
