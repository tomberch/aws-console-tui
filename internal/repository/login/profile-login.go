package login

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/rs/zerolog/log"
	"github.com/tomberch/aws-console-tui/internal/service/repository"
)

type ProfileLogin struct {
	BaseLogin
	profile string
}

func NewProfileRepository(profile string) repository.LoginRepository {
	return &ProfileLogin{profile: profile}
}

func (l *ProfileLogin) Login() aws.Config {
	config, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile(l.profile))

	if err != nil {
		log.Fatal().
			Err(err).
			Msgf("failed to load configuration")
	}

	l.checkConfig(config)
	return config
}
