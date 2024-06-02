package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		l int
		r int
	)
	flag.IntVar(&l, "l", 0, "left")
	flag.IntVar(&r, "r", 0, "right")
	flag.Parse()

	fmt.Println(l + r)
}
