package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/qnib/metahub/pkg/tooling"
)

var (
	version  = flag.Bool("version", false, "print version")
	username = flag.String("user", "metahub", "The username to login (default: metahub)")
	region   = flag.String("aws-region", "us-east-1", "AWS REGION")
	getPw    = flag.Bool("get-pass", false, "fetch password from SSM")
	getUser  = flag.Bool("get-user", false, "generate metahub login-user")
)

func main() {
	flag.Parse()
	if *version {
		fmt.Println(`v0.2.6`)
		os.Exit(0)
	}
	if !*getPw && !*getUser {
		fmt.Println("User either -get-pass or -get-user")
		os.Exit(0)
	}

	if *getUser {
		md, err := tooling.GetMetaData()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s-%s", *username, md.GetMetahubTypename())
		os.Exit(0)
	}
	if *getPw {
		passwd, err := tooling.GetSSMPassword(*region, "/metahub/password")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", passwd)
		os.Exit(0)
	}
}
