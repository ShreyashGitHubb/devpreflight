package cmd

import (
	"fmt"
	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	Version   = "0.1.0"
	BuildDate = "unknown"
	GitCommit = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Display version, build date, and commit information for DevPreflight.`,
	Example: `  # Show version information
  devpreflight version`,
	Run: func(cmd *cobra.Command, args []string) {
		printBanner()
		
		cyan := color.New(color.FgCyan).SprintFunc()
		white := color.New(color.FgWhite).SprintFunc()
		
		fmt.Printf("%s %s\n", cyan("Version:"), white(Version))
		fmt.Printf("%s %s\n", cyan("Build Date:"), white(BuildDate))
		fmt.Printf("%s %s\n", cyan("Git Commit:"), white(GitCommit))
		fmt.Printf("%s %s/%s\n", cyan("Go Version:"), white(runtime.Version()), white(runtime.GOARCH))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
