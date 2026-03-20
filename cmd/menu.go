package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "Show an interactive menu",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Println()
			fmt.Println("java-helper")
			fmt.Println("-----------")
			fmt.Println("1) Java version hints")
			fmt.Println("2) Spotless apply")
			fmt.Println("3) Verify")
			fmt.Println("4) Dependency tree")
			fmt.Println("5) Instructions about java instalation")
			fmt.Println("6) Manual instructions")
			fmt.Println("7) Exit")
			fmt.Print("Choose an option: ")

			input, err := reader.ReadString('\n')
			if err != nil {
				return err
			}

			switch strings.TrimSpace(input) {
			case "1":
				if err := javaCmd.RunE(cmd, nil); err != nil {
					return err
				}
			case "2":
				if err := fmtCmd.RunE(cmd, nil); err != nil {
					return err
				}
			case "3":
				if err := verifyCmd.RunE(cmd, nil); err != nil {
					return err
				}
			case "4":
				if err := treeCmd.RunE(cmd, nil); err != nil {
					return err
				}
			case "5":
				instructionsCmd.Run(cmd, nil)
			case "6":
				manualInstructionsCmd.Run(cmd, nil)
			case "7":
				return nil
			default:
				fmt.Println("Invalid option")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(menuCmd)
}
