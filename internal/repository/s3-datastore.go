package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func ListS3Buckets() {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		// Hard coded credentials.
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "test", SecretAccessKey: "test", SessionToken: "",
				Source: "example hard coded credentials",
			},
		}))
	if err != nil {
		log.Fatal(err)
	}

	// Credentials retrieve will be called automatically internally to the SDK
	// service clients created with the cfg value.
	creds, err := cfg.Credentials.Retrieve(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Credentials Source:", creds.Source)
	// Credentials Source: example hard coded credentials
}
