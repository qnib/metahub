package tooling

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// GetSSMPassword fetches the metahub password from SSM
func GetSSMPassword(region, path string) (passwd string, err error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(region))
	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(path),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	passwd = *param.Parameter.Value
	return
}
