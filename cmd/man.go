package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var manCmd = &cobra.Command{
	Use:   "man",
	Short: "Generate man pages",
	Long: `Generate man pages for DevPreflight commands.

Creates manual pages in the standard Unix man page format.
Install them to your system's man directory to access via 'man devpreflight'.`,
	Example: `  # Generate man pages in ./man
  devpreflight man

  # Install to system (Linux/macOS)
  sudo cp ./man/* /usr/local/share/man/man1/

  # Then use: man devpreflight`,
	RunE: func(cmd *cobra.Command, args []string) error {
		manDir := "./man"
		
		if err := os.MkdirAll(manDir, 0755); err != nil {
			return fmt.Errorf("failed to create man directory: %w", err)
		}
		
		header := &doc.GenManHeader{
			Title:   "DEVPREFLIGHT",
			Section: "1",
			Source:  "DevPreflight " + Version,
			Manual:  "DevPreflight Manual",
		}
		
		if err := doc.GenManTree(rootCmd, header, manDir); err != nil {
			return fmt.Errorf("failed to generate man pages: %w", err)
		}
		
		green := color.New(color.FgGreen).SprintFunc()
		cyan := color.New(color.FgCyan).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		
		fmt.Printf("%s Man pages generated successfully!\n", green("✓"))
		fmt.Printf("  Location: %s\n", cyan(manDir))
		fmt.Printf("\n%s To install system-wide:\n", yellow("→"))
		fmt.Printf("  sudo cp %s/*.1 /usr/local/share/man/man1/\n", manDir)
		fmt.Printf("  man devpreflight\n")
		
		return nil
	},
}

func init() {
	rootCmd.AddCommand(manCmd)
}
