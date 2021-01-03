package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// GetUsersAttrDefs returns the definitions
func GetUsersAttrDefs() []*dynamodb.AttributeDefinition {
	return []*dynamodb.AttributeDefinition{{
		AttributeName: aws.String("Login"),
		AttributeType: aws.String("S"),
	}, {
		AttributeName: aws.String("Password"),
		AttributeType: aws.String("S"),
	}}
}

// GetUsersAttrSchemas returns the schemas
func GetUsersAttrSchemas() []*dynamodb.KeySchemaElement {
	return []*dynamodb.KeySchemaElement{{
		AttributeName: aws.String("Login"),
		KeyType:       aws.String("HASH"),
	}, {
		AttributeName: aws.String("Password"),
		KeyType:       aws.String("RANGE"),
	}}
}

// UsersItem holds the features for a particular login
type UsersItem struct {
	Login    string
	Password string
}
