package cmd

import (
        "fmt"
        "os"

        "github.com/fatih/color"
        "github.com/spf13/cobra"
        "github.com/spf13/viper"
)

const (
        ExitSuccess = 0
        ExitWarning = 1
        ExitError   = 2
)

var (
        cfgFile  string
        verbose  bool
        noColor  bool
)

var rootCmd = &cobra.Command{
        Use:   "devpreflight",
        Short: "Developer deployment preflight checks",
        Long: `DevPreflight – Developer Deployment Preflight Checks

A CLI tool that runs curated preflight checks to catch issues before 
they hit production. Validates environment parity, Dockerfile best practices,
Kubernetes manifests, observability instrumentation, and flaky tests.

Config file: .devpreflightrc.yml (auto-detected in current directory)`,
        Example: `  # Run all checks with default settings
  devpreflight check

  # Generate JSON report for CI/CD
  devpreflight ci-report --format json --output report.json

  # Fix environment variable issues
  devpreflight fix --env --yes

  # Generate documentation
  devpreflight docs

  # Show version
  devpreflight version`,
}

func Execute() error {
        err := rootCmd.Execute()
        if err != nil {
                os.Exit(ExitError)
        }
        return nil
}

func init() {
        cobra.OnInitialize(initConfig)
        rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .devpreflightrc.yml)")
        rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
        rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable colored output")
        viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func printBanner() {
        cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
        gray := color.New(color.FgHiBlack).SprintFunc()
        
        fmt.Println()
        fmt.Printf("%s %s %s\n", cyan("DevPreflight"), gray("v"+Version), gray("– Developer Deployment Preflight Checks"))
        fmt.Println()
}

func initConfig() {
        if cfgFile != "" {
                viper.SetConfigFile(cfgFile)
        } else {
                viper.SetConfigName(".devpreflightrc")
                viper.SetConfigType("yml")
                viper.AddConfigPath(".")
        }

        viper.AutomaticEnv()

        if err := viper.ReadInConfig(); err == nil {
                fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
        }
}
