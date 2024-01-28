package main

import (
	"flag"
	"fmt"
)

func main() {
	nFlag := flag.Int("n", 23, "start number")
	flag.Parse()

	fmt.Println(*nFlag)
}
