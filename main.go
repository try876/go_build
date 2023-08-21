package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
)

var (
	ver = flag.Bool("v", false, "")
)

//go:embed version.txt
var version string

func main() {
	flag.Parse()
	if *ver {
		fmt.Println("print version:", version)
		os.Exit(0)
	}

	fmt.Println("hello world:", version)
}
