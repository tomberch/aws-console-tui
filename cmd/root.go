package cmd

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
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
	// initK9sFlags()
	// initK8sFlags()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Panic().Err(err)
	}
}

func run(cmd *cobra.Command, args []string) error {
	fmt.Println("Hello, I am here")
	return nil
}
