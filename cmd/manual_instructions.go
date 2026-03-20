package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var manualInstructionsCmd = &cobra.Command{
	Use:   "manual-instructions",
	Short: "Show the manual Maven commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Manual instructions")
		fmt.Println("-------------------")
		fmt.Println()
		fmt.Println("Check Java version")
		fmt.Println("  mvn help:effective-pom")
		fmt.Println()
		fmt.Println("Format the project")
		fmt.Println("  mvn -U spotless:apply")
		fmt.Println()
		fmt.Println("Compile and run tests")
		fmt.Println("  mvn verify -U --batch-mode -Dbasepom.check.skip-dependency-versions-check=true -Dorg.slf4j.simpleLogger.log.org.apache.maven.cli.transfer.Slf4jMavenTransferListener=warn")
		fmt.Println()
		fmt.Println("Show the dependency tree")
		fmt.Println("  mvn dependency:tree")
	},
}

func init() {
	rootCmd.AddCommand(manualInstructionsCmd)
}
