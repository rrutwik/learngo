package main

import (
	"os"
	"time"

	"example.com/hello/maths"
)

func main() {
	t := time.Now()
	maths.SVGWriter(os.Stdout, t)
}
