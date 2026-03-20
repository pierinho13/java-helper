package cmd

import (
	"github.com/spf13/cobra"

	"github.com/pierinho13/java-helper/internal/runner"
)

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Run mvn dependency:tree",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runner.Run("mvn", "dependency:tree")
	},
}

func init() {
	rootCmd.AddCommand(treeCmd)
}
