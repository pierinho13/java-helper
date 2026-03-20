package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "java-helper",
	Short: "Small CLI to run common Maven tasks for local Java projects",
	Long:  "java-helper provides shortcuts for common Maven commands used in Java projects.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return menuCmd.RunE(cmd, args)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
