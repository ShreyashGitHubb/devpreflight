package cmd

import (
	"fmt"
	"os"

	"github.com/devpreflight/devpreflight/internal/checks"
	"github.com/devpreflight/devpreflight/internal/config"
	"github.com/devpreflight/devpreflight/internal/reporter"
)

func runChecksAndReport(cfg *config.Config, format, outputFile string) int {
	results := checks.RunAllChecks(cfg)
	
	rep := selectReporter(format)
	content := rep.Report(results)
	
	if outputFile != "" {
		if err := os.WriteFile(outputFile, []byte(content), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to %s: %v\n", outputFile, err)
			return 1
		}
	} else {
		fmt.Print(content)
	}
	
	return determineExitCode(results)
}

func selectReporter(format string) reporter.Reporter {
	switch format {
	case "json":
		return reporter.NewJSONReporter()
	case "markdown":
		return reporter.NewMarkdownReporter()
	default:
		return reporter.NewConsoleReporter()
	}
}

func determineExitCode(results []checks.CheckResult) int {
	hasFailures := false
	hasWarnings := false
	
	for _, result := range results {
		if result.Status == checks.StatusFail {
			hasFailures = true
		} else if result.Status == checks.StatusWarn {
			hasWarnings = true
		}
	}
	
	if hasFailures {
		return 2
	} else if hasWarnings {
		return 1
	}
	return 0
}
