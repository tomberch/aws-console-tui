package repository

import "github.com/aws/aws-sdk-go-v2/aws"

type LoginRepository interface {
	Login() aws.Config
}
