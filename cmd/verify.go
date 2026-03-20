package cmd

import (
	"github.com/spf13/cobra"

	"github.com/pierinho13/java-helper/internal/runner"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Run mvn verify with the required local flags",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runner.Run(
			"mvn",
			"verify",
			"-U",
			"--batch-mode",
			"-Dbasepom.check.skip-dependency-versions-check=true",
			"-Dorg.slf4j.simpleLogger.log.org.apache.maven.cli.transfer.Slf4jMavenTransferListener=warn",
		)
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}
