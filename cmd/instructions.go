package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func detectHomebrewOptPath() string {
	candidates := []string{
		"/opt/homebrew/opt",
		"/usr/local/opt",
	}

	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	return "<homebrew-opt-path>"
}

var instructionsCmd = &cobra.Command{
	Use:   "instructions",
	Short: "Show Java and Maven installation instructions",
	Run: func(cmd *cobra.Command, args []string) {
		brewOpt := detectHomebrewOptPath()

		fmt.Println("Installation instructions")
		fmt.Println("-------------------------")
		fmt.Println()
		fmt.Println("Install Maven")
		fmt.Println("  brew install maven")
		fmt.Println()
		fmt.Println("Install jenv")
		fmt.Println("  brew install jenv")
		fmt.Println()
		fmt.Println("Install Java 17")
		fmt.Println("  brew install openjdk@17")
		fmt.Printf("  jenv add %s/openjdk@17\n", brewOpt)
		fmt.Println()
		fmt.Println("Install Java 21")
		fmt.Println("  brew install openjdk@21")
		fmt.Printf("  jenv add %s/openjdk@21\n", brewOpt)
		fmt.Println()
		fmt.Println("Install Java 25")
		fmt.Println("  brew install openjdk@25")
		fmt.Printf("  jenv add %s/openjdk\n", brewOpt)
		fmt.Println()
		fmt.Println("Change Java version with jenv")
		fmt.Println("  jenv global <version>")
		fmt.Println("  example: jenv global 17")
	},
}

func init() {
	rootCmd.AddCommand(instructionsCmd)
}
