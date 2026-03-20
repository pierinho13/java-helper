package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "doodle-java",
	Short: "Small CLI to run common Maven tasks for Doodle Java projects",
	Long:  "doodle-java provides shortcuts for common Maven commands used in Java projects.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
