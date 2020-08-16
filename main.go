package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jcbritobr/img2ascii/encoder"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	filename := flag.String("f", "sample.jpg", "-f <filename>")
	scale := flag.Uint("s", 70, "-s <scale>")
	flag.Parse()

	file, err := os.OpenFile(*filename, os.O_RDONLY, 0644)
	checkError(err)

	result, err := encoder.Encode(file, *scale)
	checkError(err)
	fmt.Print(result)
}
