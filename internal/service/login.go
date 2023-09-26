package service

import (
	"github.com/tomberch/aws-console-tui/internal/appconfig"
	"github.com/tomberch/aws-console-tui/internal/repository/login"
	"github.com/tomberch/aws-console-tui/internal/service/repository"
)

func Login(profile string) {
	loginRepository := createLoginRepository(profile)
	config := loginRepository.Login()
	appconfig.GetInstance().SetAwsConfig(config)
}

func createLoginRepository(profile string) repository.LoginRepository {
	if profile == "" {
		return login.NewDefaultLogin()
	}

	return login.NewProfileRepository(profile)
}
