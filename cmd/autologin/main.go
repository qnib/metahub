package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	version  = flag.Bool("version", false, "print version")
	username = flag.String("user", "metahub", "The username to login (default: metahub)")
)

func main() {
	flag.Parse()
	if *version {
		fmt.Println(`v0.2.5`)
		os.Exit(0)
	}
}
