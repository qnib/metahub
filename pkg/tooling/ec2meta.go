package tooling

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/klauspost/cpuid"
)

type EC2Meta struct {
	InstanceType   string
	InstanceSize   string
	HyperThreading string
}

// GetMetaData fetches the metadata of an instance
func GetMetaData() (md EC2Meta, err error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := ec2metadata.New(sess)
	instType, err := svc.GetMetadata("instance-type")
	iTypeSlice := strings.Split(instType, ".")
	md.InstanceType = iTypeSlice[0]
	r, _ := regexp.Compile("[0-9]*xl")
	md.InstanceSize = r.FindString(iTypeSlice[1])
	switch cpuid.CPU.ThreadsPerCore {
	case 1:
		md.HyperThreading = "off"
	case 2:
		md.HyperThreading = "on"
	}
	return
}

// GetMetahubFeatures returns the features
func (md *EC2Meta) GetMetahubFeatures() []string {
	return []string{
		fmt.Sprintf("instType:%s", md.InstanceType),
		fmt.Sprintf("instSize:%s", md.InstanceSize),
		fmt.Sprintf("instHT:%s", md.HyperThreading),
	}
}

// GetMetahubTypename returns the typename to login to metahub
func (md *EC2Meta) GetMetahubTypename() string {
	res := []string{fmt.Sprintf("%s%s", md.InstanceType, md.InstanceSize)}
	if md.HyperThreading == "on" {
		res = append(res, "ht")
	}
	return strings.Join(res, "-")
}
