package main

import (
	"log"

	w "github.com/lepeico/withstyle"
)

var (
	debug   = w.WithStyles(w.BgWhite, w.Black, w.Italic)
	error   = w.WithStyles(w.BgRed, w.Black, w.Bold)
	warning = w.WithStyles(w.BgOrange, w.White, w.Underline)
	success = w.WithStyles(w.BgGreen, w.White, w.Italic)
	info    = w.WithStyles(w.BgBlue, w.Black, w.Italic)
	verbose = w.WithStyles(w.BgDeepPurple, w.White)
	fatal   = w.WithStyles(w.BgBlack, w.Red, w.Bold)
)

func main() {
	log.Println(debug("Okey. This is a debug message!"))
	log.Println(error("Oops! This is an error message!"))
	log.Println(warning("Hm.. This is a warning message!"))
	log.Println(success("Yay! This is a success message!"))
	log.Println(info("Hi! Just an info message!"))
	log.Println(verbose("So.. This is a verbose message!"))
	log.Println(fatal("%*#$! This is a fatal message!"))
}
