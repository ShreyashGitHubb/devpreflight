package cmd

import (
        "os"

        "github.com/devpreflight/devpreflight/internal/config"
        "github.com/spf13/cobra"
        "github.com/spf13/viper"
)

var checkCmd = &cobra.Command{
        Use:   "check",
        Short: "Run all preflight checks",
        Long: `Run all enabled preflight checks and report results.

Checks performed:
  • Environment parity (.env vs .env.example)
  • Dockerfile best practices (no :latest tags, multi-stage builds)
  • Kubernetes manifest validation (deprecated APIs, required fields)
  • Observability instrumentation (OTEL, Sentry, Prometheus)
  • Flaky test detection (runs tests multiple times)

The tool returns exit code 0 for success, 1 for warnings, and 2 for failures.
This makes it perfect for CI/CD pipelines that need to fail on critical issues.`,
        Example: `  # Run all checks with console output
  devpreflight check

  # Run checks and save JSON report
  devpreflight check --format json --output report.json

  # Run checks with Markdown output
  devpreflight check --format markdown

  # Use custom config file
  devpreflight check --config ./custom-config.yml`,
        RunE: func(cmd *cobra.Command, args []string) error {
                if verbose {
                        printBanner()
                }
                
                cfg := config.LoadConfig()
                
                format, _ := cmd.Flags().GetString("format")
                if format == "" {
                        format = cfg.Report.Format
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
        rootCmd.AddCommand(checkCmd)
        checkCmd.Flags().String("format", "", "Output format: console, json, markdown (default from config)")
        checkCmd.Flags().String("output", "", "Output file (default: stdout)")
        viper.BindPFlag("report.format", checkCmd.Flags().Lookup("format"))
        viper.BindPFlag("report.output", checkCmd.Flags().Lookup("output"))
}
