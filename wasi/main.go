package main

import (
	"flag"
)

func main() {
	var (
		l int
		r int
	)
	flag.IntVar(&l, "l", 0, "left")
	flag.IntVar(&r, "r", 0, "right")
	flag.Parse()

	println(l + r)
}
