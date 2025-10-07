package cmd

import (
        "os"

        "github.com/devpreflight/devpreflight/internal/config"
        "github.com/spf13/cobra"
)

var ciReportCmd = &cobra.Command{
        Use:   "ci-report",
        Short: "Generate CI-friendly reports",
        Long: `Generate reports in JSON or Markdown format for CI/CD integration.

This command is optimized for continuous integration pipelines:
  • Defaults to JSON format (machine-parseable)
  • Includes exit codes in output
  • Can write to files for artifact storage
  • Markdown format creates PR-ready comments

The report includes summary statistics, check details, and remediation steps.
Exit codes: 0 (success), 1 (warnings), 2 (failures).`,
        Example: `  # Generate JSON report to stdout
  devpreflight ci-report

  # Generate JSON report to file
  devpreflight ci-report --output report.json

  # Generate Markdown report for PR comments
  devpreflight ci-report --format markdown --output pr-comment.md

  # Use in GitHub Actions
  devpreflight ci-report --format markdown --output $GITHUB_STEP_SUMMARY`,
        RunE: func(cmd *cobra.Command, args []string) error {
                if verbose {
                        printBanner()
                }
                
                cfg := config.LoadConfig()
                
                format, _ := cmd.Flags().GetString("format")
                if format == "" {
                        // Default to json for CI if not specified in config
                        if cfg.Report.Format == "" || cfg.Report.Format == "console" {
                                format = "json"
                        } else {
                                format = cfg.Report.Format
                        }
                }
                
                output, _ := cmd.Flags().GetString("output")
                if output == "" {
                        output = cfg.Report.Output
                }
                
                exitCode := runChecksAndReport(cfg, format, output)
                os.Exit(exitCode)
                return nil
        },
}

func init() {
        rootCmd.AddCommand(ciReportCmd)
        ciReportCmd.Flags().String("format", "", "Output format: json, markdown (default: json for CI, or from config)")
        ciReportCmd.Flags().String("output", "", "Output file (default: stdout)")
}
