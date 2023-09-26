package appconfig

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/rs/zerolog/log"
)

type appConfig struct {
	awsConfig aws.Config
}

var lock = &sync.Mutex{}
var singleConfig *appConfig

func GetInstance() *appConfig {
	if singleConfig == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleConfig == nil {
			log.Debug().Msg("Creating app configuration")
			singleConfig = &appConfig{}
		} else {
			log.Debug().Msg("Accessing existing app configuration instance")
		}
	} else {
		log.Debug().Msg("Accessing existing app configuration instance")
	}

	return singleConfig
}

func (appConfig *appConfig) GetAwsConfig() aws.Config {
	return appConfig.awsConfig
}

func (appConfig *appConfig) SetAwsConfig(awsConfig aws.Config) {
	appConfig.awsConfig = awsConfig
}
