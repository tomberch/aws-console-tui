package login

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/rs/zerolog/log"
	"github.com/tomberch/aws-console-tui/internal/service/repository"
)

type DefaultLogin struct {
	BaseLogin
}

func NewDefaultLogin() repository.LoginRepository {
	return &DefaultLogin{}
}

func (l *DefaultLogin) Login() aws.Config {
	config, err := config.LoadDefaultConfig(
		context.TODO())

	if err != nil {
		log.Fatal().
			Err(err).
			Msgf("failed to load configuration")
	}

	l.checkConfig(config)
	return config
}
