package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/qnib/metahub/pkg/tooling"
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
	md, err := tooling.GetMetaData()
	if err != nil {
		panic(err)
	}
	log.Println(md.GetMetahubFeatures())
	log.Println(md.GetMetahubTypename())
}
