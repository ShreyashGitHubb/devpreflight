package cmd

import (
        "fmt"

        "github.com/devpreflight/devpreflight/internal/config"
        "github.com/devpreflight/devpreflight/internal/fixer"
        "github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
        Use:   "fix",
        Short: "Auto-fix issues where possible",
        Long: `Attempt to automatically fix detected issues.

Currently supported fixes:
  --env    Add missing environment variables to .env with __REPLACE_ME__ placeholders

The fix command will never populate actual values automatically for security reasons.
You must manually replace placeholder values with real credentials.

Other checks (Dockerfile, K8s, observability, flaky tests) provide recommendations
but do not have automated fixes yet. These require manual intervention.`,
        Example: `  # Fix environment variable parity (interactive)
  devpreflight fix --env

  # Fix environment variable parity (auto-confirm)
  devpreflight fix --env --yes

  # Show what can be fixed
  devpreflight fix`,
        RunE: func(cmd *cobra.Command, args []string) error {
                if verbose {
                        printBanner()
                }
                
                cfg := config.LoadConfig()
                
                env, _ := cmd.Flags().GetBool("env")
                yes, _ := cmd.Flags().GetBool("yes")
                
                if env {
                        return fixer.FixEnvParity(cfg, yes)
                }
                
                fmt.Println("Specify what to fix:")
                fmt.Println("  --env    Fix environment variable parity")
                fmt.Println("\nUse 'devpreflight fix --help' for more information")
                return nil
        },
}

func init() {
        rootCmd.AddCommand(fixCmd)
        fixCmd.Flags().Bool("env", false, "Fix environment variable parity")
        fixCmd.Flags().Bool("yes", false, "Auto-confirm all fixes")
}
