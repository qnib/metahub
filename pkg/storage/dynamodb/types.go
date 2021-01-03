package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// GetTypesAttrDefs returns the definitions
func GetTypesAttrDefs() []*dynamodb.AttributeDefinition {
	return []*dynamodb.AttributeDefinition{{
		AttributeName: aws.String("Type"),
		AttributeType: aws.String("S"),
	}, {
		AttributeName: aws.String("Features"),
		AttributeType: aws.String("S"),
	}}
}

// GetTypesAttrSchemas returns the schemas
func GetTypesAttrSchemas() []*dynamodb.KeySchemaElement {
	return []*dynamodb.KeySchemaElement{{
		AttributeName: aws.String("Type"),
		KeyType:       aws.String("HASH"),
	}, {
		AttributeName: aws.String("Features"),
		KeyType:       aws.String("RANGE"),
	}}
}

// TypesItem holds the features for a particular type of client
type TypeItem struct {
	Type     string
	Features string
}
