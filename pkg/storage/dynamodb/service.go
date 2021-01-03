package dynamodb

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/qnib/metahub/pkg/storage"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	svc    *dynamodb.DynamoDB
	dbSync sync.Mutex
)

const (
	mhDbTablePrefix = "metahub-v1"
)

func init() {
	err := setupDB()
	if err != nil {
		log.Fatal(err.Error())
	}
}

// NewService returns a new storage.Service for boltdb
func NewService() storage.Service {
	return &service{}
}

type service struct{}

func (s *service) MachineTypeService(ctx context.Context) (storage.MachineTypeService, error) {
	return &machineTypeService{
		ctx: ctx,
	}, nil
}

func (s *service) AccessTokenService(ctx context.Context) (storage.AccessTokenService, error) {
	return &accessTokenService{
		ctx: ctx,
	}, nil
}

func (s *service) AccountService(ctx context.Context) (storage.AccountService, error) {
	return &accountService{
		ctx: ctx,
	}, nil
}

func setupDB() (err error) {
	dbSync.Lock()
	defer dbSync.Unlock()
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc = dynamodb.New(sess)
	if !mhTableExists(svc, fmt.Sprintf("%s_types", mhDbTablePrefix)) {
		err = mhTableTypesCreate(svc)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	if !mhTableExists(svc, fmt.Sprintf("%s_users", mhDbTablePrefix)) {
		err = mhTableUsersCreate(svc)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = mhTableInit(svc)
	}

	fmt.Println("DB Setup Done")
	return
}

func mhTableExists(db *dynamodb.DynamoDB, mhDbTable string) bool {
	input := &dynamodb.ListTablesInput{}
	for {
		// Get the list of tables
		result, err := db.ListTables(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					log.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					log.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
		}
		for _, n := range result.TableNames {
			// if we already find the table we are looking for - awesome
			// -> otherwise rember we need to create it.
			log.Printf("Table found: %s\n", *n)
			if *n == mhDbTable {
				return true
			}

		}
		// assign the last read tablename as the start for our next call to the ListTables function
		// the maximum number of table names returned in a call is 100 (default), which requires us to make
		// multiple calls to the ListTables function to retrieve all table names
		input.ExclusiveStartTableName = result.LastEvaluatedTableName
		if result.LastEvaluatedTableName == nil {
			break
		}
	}
	return false
}

func mhTableInit(db *dynamodb.DynamoDB) (err error) {
	return
}

func mhTableUsersCreate(db *dynamodb.DynamoDB) (err error) {
	mhDbTable := fmt.Sprintf("%s_users", mhDbTablePrefix)
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: GetUsersAttrDefs(),
		KeySchema:            GetUsersAttrSchemas(),
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(mhDbTable),
	}

	_, err = db.CreateTable(input)
	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
	}
	fmt.Println("Created the table", mhDbTable)
	return
}

func mhTableTypesCreate(db *dynamodb.DynamoDB) (err error) {
	mhDbTable := fmt.Sprintf("%s_types", mhDbTablePrefix)
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: GetTypesAttrDefs(),
		KeySchema:            GetTypesAttrSchemas(),
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(mhDbTable),
	}

	_, err = db.CreateTable(input)
	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
	}
	fmt.Println("Created the table", mhDbTable)
	return
}

func mhTableUserScan(db *dynamodb.DynamoDB, tableName, usern string) (user UsersItem, err error) {
	log.Printf("Search for username '%s' in '%s'", usern, tableName)
	var queryInput = &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		KeyConditions: map[string]*dynamodb.Condition{
			"Login": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(usern),
					},
				},
			},
		},
	}
	result, err := db.Query(queryInput)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	userObj := []UsersItem{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &userObj)

	switch len(userObj) {
	case 0:
		err = fmt.Errorf("Could not find user: '%s'", usern)
	case 1:
		user = userObj[0]
	default:
		err = fmt.Errorf("Found multiple users for '%s'? WTF?", usern)
	}
	return
}

func mhTableTypeScan(db *dynamodb.DynamoDB, tableName, typen string) (typeItem TypeItem, err error) {
	log.Printf("Search for type '%s' in '%s'", typen, tableName)
	var queryInput = &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		KeyConditions: map[string]*dynamodb.Condition{
			"Type": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(typen),
					},
				},
			},
		},
	}
	result, err := db.Query(queryInput)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	typeObj := []TypeItem{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &typeObj)

	switch len(typeObj) {
	case 0:
		err = fmt.Errorf("Could not find type: '%s'", typen)
	case 1:
		typeItem = typeObj[0]
	default:
		err = fmt.Errorf("Found multiple types for '%s'? WTF?", typen)
	}
	return
}
