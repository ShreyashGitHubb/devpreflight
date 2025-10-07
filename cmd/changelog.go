package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var changelogCmd = &cobra.Command{
	Use:   "changelog",
	Short: "Display the changelog",
	Long: `Display the DevPreflight changelog showing version history and changes.

The changelog follows the Keep a Changelog format and shows all notable
changes for each version of DevPreflight.`,
	Example: `  # Display the changelog
  devpreflight changelog

  # View in a pager
  devpreflight changelog | less`,
	Run: func(cmd *cobra.Command, args []string) {
		changelogPath := "CHANGELOG.md"
		
		content, err := os.ReadFile(changelogPath)
		if err != nil {
			// If CHANGELOG.md doesn't exist, show embedded changelog
			yellow := color.New(color.FgYellow).SprintFunc()
			fmt.Printf("%s CHANGELOG.md not found, showing current version info:\n\n", yellow("⚠"))
			printBanner()
			fmt.Printf("Version: %s\n", Version)
			fmt.Printf("Build Date: %s\n", BuildDate)
			fmt.Printf("\nFor full changelog, see: https://github.com/devpreflight/devpreflight/blob/main/CHANGELOG.md\n")
			return
		}
		
		fmt.Print(string(content))
	},
}

func init() {
	rootCmd.AddCommand(changelogCmd)
}
