package main

import (
	"flag"
	"fmt"
)

func main() {
	nFlag := flag.Int("n", 23, "start number")
	flag.Parse()

	n := *nFlag
	fmt.Println(n)

	if n%2 == 1 {
		n = n*3 + 1
	} else {
		n /= 2
	}
	fmt.Println(n)
}
