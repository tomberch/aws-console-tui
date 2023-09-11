package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	var short bool

	command := cobra.Command{
		Use:   "version",
		Short: "Print version/build info",
		Long:  "Print version/build information",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion(short)
		},
	}

	command.PersistentFlags().BoolVarP(&short, "short", "s", false, "Prints aws-console-tui version info in short format")

	return &command
}

func printVersion(short bool) {
	if short {
		fmt.Println("Version: 0.0.1")
	} else {
		fmt.Println("Version: 0.0.1")
		fmt.Println("Build: Build Hash")
	}
}
