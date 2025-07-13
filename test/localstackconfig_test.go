package test

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	cfg "github.com/vmware/vmware-go-kcl/clientlibrary/config"
)

// configureLocalStackEndpoints configures KCL config to use LocalStack endpoints if environment variables are set
func configureLocalStackEndpoints(kclConfig *cfg.KinesisClientLibConfiguration) *cfg.KinesisClientLibConfiguration {
	// Configure Kinesis endpoint if specified
	if kinesisEndpoint := os.Getenv("KINESIS_ENDPOINT"); kinesisEndpoint != "" {
		kclConfig = kclConfig.WithKinesisEndpoint(kinesisEndpoint)
	}

	// Configure DynamoDB endpoint if specified
	if dynamoDBEndpoint := os.Getenv("DYNAMODB_ENDPOINT"); dynamoDBEndpoint != "" {
		kclConfig = kclConfig.WithDynamoDBEndpoint(dynamoDBEndpoint)
	}

	// Use local credentials if we're using LocalStack endpoints
	if os.Getenv("KINESIS_ENDPOINT") != "" || os.Getenv("DYNAMODB_ENDPOINT") != "" {
		// LocalStack accepts any credentials
		localCreds := credentials.NewStaticCredentials("test", "test", "")
		kclConfig.KinesisCredentials = localCreds
		kclConfig.DynamoDBCredentials = localCreds
	}

	return kclConfig
}
