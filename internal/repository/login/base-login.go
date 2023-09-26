package login

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type BaseLogin struct {
}

func (b *BaseLogin) checkConfig(config aws.Config) {
	creds, err := config.Credentials.Retrieve(context.TODO())
	if err != nil {
		log.Fatal().
			Err(err).
			Msgf("Invalid configuration, cannot retrieve credentials")
	}

	if log.Logger.GetLevel() <= zerolog.DebugLevel {
		log.Debug().
			Str("AccessKeyId", creds.AccessKeyID).
			Str("Source", creds.Source).
			Str("Region", config.Region).
			Msg("Credentials Info")
	}
}
