package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/tomberch/aws-console-tui/internal/service"
)

const (
	appName      = "aws-console-tui"
	shortAppDesc = "An AWS console within your terminal."
	longAppDesc  = "aws-console-tui provides a AWS console directly within your terminal to create, browse, or edit your services."
)

var (
	version, commitHash = "none", "none" // will be replaced during build with linker flags

	rootCmd = &cobra.Command{
		Use:   appName,
		Short: shortAppDesc,
		Long:  longAppDesc,
		RunE:  run,
	}
)

func init() {
	rootCmd.AddCommand(versionCmd())

	rootCmd.PersistentFlags().String("profile", "", "AWS profile to be used to create a configuration.")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Panic().Err(err)
	}
}

func run(cmd *cobra.Command, args []string) error {
	profile, err := cmd.Flags().GetString("profile")

	if err != nil {
		log.Fatal().
			Err(err).
			Msgf("Unable to fetch flag 'profile'")
	}

	service.Login(profile)

	return nil
}
