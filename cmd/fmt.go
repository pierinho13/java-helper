package cmd

import (
	"github.com/spf13/cobra"

	"github.com/pierinho13/doodle-java/internal/runner"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Run spotless apply",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runner.Run("mvn", "-U", "spotless:apply")
	},
}

func init() {
	rootCmd.AddCommand(fmtCmd)
}
