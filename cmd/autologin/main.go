package main

import (
	"flag"
	"fmt"
	"os"

	"log"

	"github.com/qnib/metahub/pkg/tooling"
)

var (
	version     = flag.Bool("version", false, "print version")
	username    = flag.String("user", "metahub", "The username to login (default: metahub)")
	typename    = flag.String("type", "", "Define the machine type, will be generated based on instance info otherwise")
	region      = flag.String("aws-region", "", "AWS REGION")
	regname     = flag.String("registry", "mh.qnib.org", "Metahub registry name")
	getPw       = flag.Bool("get-pass", false, "fetch password from SSM")
	getUser     = flag.Bool("get-user", false, "generate metahub login-user")
	dockerLogin = flag.Bool("docker-login", false, "Use the docker-client to login directly (e.g. /var/run/docker.sock)")
)

func main() {
	flag.Parse()
	if *version {
		fmt.Println(`v0.2.9`)
		os.Exit(0)
	}
	awsRegion := ""
	switch {
	case *region == "" && os.Getenv("AWS_REGION") != "":
		awsRegion = os.Getenv("AWS_REGION")
	case *region != "" && os.Getenv("AWS_REGION") != "":
		log.Printf("--aws-region is set while AWS_REGION is also present; we'll go with the CLI argument")
		awsRegion = *region
	case *region == "" && os.Getenv("AWS_REGION") == "" && !*dockerLogin:
		log.Printf("--aws-region and $AWS_REGION are both empty")
	default:
		awsRegion = *region
	}
	if !*getPw && !*getUser && !*dockerLogin {
		fmt.Println("Use: -get-pass, -docker-login or -get-user")
		os.Exit(0)
	}
	if *dockerLogin {
		md, err := tooling.GetMetaData()
		if err != nil {
			panic(err)
		}
		uname := fmt.Sprintf("%s-%s", *username, md.GetMetahubTypename(*typename))
		passwd, err := tooling.GetSSMPassword(awsRegion, "/metahub/password")
		if err != nil {
			panic(err)
		}
		err = tooling.DockerLogin(string(passwd), uname, *regname)
		if err != nil {
			log.Print(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}
	if *getUser {
		md, err := tooling.GetMetaData()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s-%s\n", *username, md.GetMetahubTypename(*typename))
		os.Exit(0)
	}
	if *getPw {
		passwd, err := tooling.GetSSMPassword(awsRegion, "/metahub/password")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", passwd)
		os.Exit(0)
	}
}
