package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate documentation",
	Long: `Generate Markdown documentation for all DevPreflight commands.

Creates a /docs directory with individual markdown files for each command.
These can be used for websites, wikis, or version control documentation.`,
	Example: `  # Generate documentation in ./docs
  devpreflight docs

  # The following files will be created:
  # - docs/devpreflight.md
  # - docs/devpreflight_check.md
  # - docs/devpreflight_fix.md
  # - docs/devpreflight_ci-report.md
  # - docs/devpreflight_completion.md
  # - docs/devpreflight_version.md`,
	RunE: func(cmd *cobra.Command, args []string) error {
		docsDir := "./docs"
		
		if err := os.MkdirAll(docsDir, 0755); err != nil {
			return fmt.Errorf("failed to create docs directory: %w", err)
		}
		
		if err := doc.GenMarkdownTree(rootCmd, docsDir); err != nil {
			return fmt.Errorf("failed to generate documentation: %w", err)
		}
		
		green := color.New(color.FgGreen).SprintFunc()
		cyan := color.New(color.FgCyan).SprintFunc()
		
		fmt.Printf("%s Documentation generated successfully!\n", green("✓"))
		fmt.Printf("  Location: %s\n", cyan(docsDir))
		fmt.Printf("  Files: %s\n", cyan("devpreflight*.md"))
		
		return nil
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
