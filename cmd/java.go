package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/pierinho13/doodle-java/internal/runner"
)

var javaCmd = &cobra.Command{
	Use:   "java",
	Short: "Inspect Java-related Maven config from effective-pom",
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := runner.RunAndCapture("mvn", "help:effective-pom")
		if err != nil {
			return err
		}

		patterns := []string{
			"<java.version>",
			"<maven.compiler.source>",
			"<maven.compiler.target>",
			"<maven.compiler.release>",
			"<release>",
			"<source>",
			"<target>",
		}

		scanner := bufio.NewScanner(bytes.NewReader(out))
		found := false

		fmt.Println("Relevant Java-related lines from effective-pom:")
		fmt.Println("------------------------------------------------")

		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			for _, p := range patterns {
				if strings.Contains(line, p) {
					fmt.Println(line)
					found = true
					break
				}
			}
		}

		fmt.Println("------------------------------------------------")

		if !found {
			fmt.Println("No obvious Java version lines found in effective-pom.")
		}

		fmt.Println("\nLocal Java:")
		_ = runner.Run("java", "-version")

		fmt.Println("\nLocal Maven:")
		return runner.Run("mvn", "-version")
	},
}

func init() {
	rootCmd.AddCommand(javaCmd)
}
