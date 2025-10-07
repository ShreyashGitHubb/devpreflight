package reporter

import (
        "fmt"
        "strings"

        "github.com/devpreflight/devpreflight/internal/checks"
        "github.com/fatih/color"
)

type ConsoleReporter struct{}

func NewConsoleReporter() *ConsoleReporter {
        return &ConsoleReporter{}
}

func (r *ConsoleReporter) Report(results []checks.CheckResult) string {
        var output strings.Builder
        
        // Header
        cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
        output.WriteString("\n")
        output.WriteString(cyan("DevPreflight Check Results"))
        output.WriteString("\n")
        output.WriteString(cyan("==========================="))
        output.WriteString("\n\n")
        
        failCount := 0
        warnCount := 0
        okCount := 0
        
        for _, result := range results {
                var symbol string
                var colorFunc func(a ...interface{}) string
                
                switch result.Status {
                case checks.StatusOK:
                        symbol = "✔"
                        colorFunc = color.New(color.FgGreen).SprintFunc()
                        okCount++
                case checks.StatusWarn:
                        symbol = "⚠"
                        colorFunc = color.New(color.FgYellow).SprintFunc()
                        warnCount++
                case checks.StatusFail:
                        symbol = "✖"
                        colorFunc = color.New(color.FgRed).SprintFunc()
                        failCount++
                case checks.StatusSkip:
                        symbol = "⏭"
                        colorFunc = color.New(color.FgHiBlack).SprintFunc()
                }
                
                gray := color.New(color.FgHiBlack).SprintFunc()
                
                output.WriteString(fmt.Sprintf("%s %s: %s\n", 
                        colorFunc(symbol), 
                        gray(result.Name), 
                        result.Message))
                
                if result.Details != "" && result.Status != checks.StatusOK && result.Status != checks.StatusSkip {
                        // Indent details
                        lines := strings.Split(result.Details, "\n")
                        for _, line := range lines {
                                if line != "" {
                                        output.WriteString(fmt.Sprintf("  %s\n", gray(line)))
                                }
                        }
                }
        }
        
        // Summary
        output.WriteString("\n")
        green := color.New(color.FgGreen).SprintFunc()
        yellow := color.New(color.FgYellow).SprintFunc()
        red := color.New(color.FgRed).SprintFunc()
        
        output.WriteString(fmt.Sprintf("Summary: %s passed, %s warnings, %s failures\n", 
                green(fmt.Sprintf("%d", okCount)),
                yellow(fmt.Sprintf("%d", warnCount)),
                red(fmt.Sprintf("%d", failCount))))
        
        // Remediation hints
        if failCount > 0 {
                cyan := color.New(color.FgCyan).SprintFunc()
                output.WriteString("\n")
                output.WriteString(fmt.Sprintf("%s Run 'devpreflight fix --env' to fix environment issues\n", cyan("→")))
                output.WriteString(fmt.Sprintf("%s Or 'devpreflight ci-report --format markdown' for detailed report\n", cyan("→")))
        }
        
        return output.String()
}
