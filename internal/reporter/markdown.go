package reporter

import (
	"fmt"
	"strings"

	"github.com/devpreflight/devpreflight/internal/checks"
)

type MarkdownReporter struct{}

func NewMarkdownReporter() *MarkdownReporter {
	return &MarkdownReporter{}
}

func (r *MarkdownReporter) Report(results []checks.CheckResult) string {
	var output strings.Builder
	
	output.WriteString("# DevPreflight Check Results\n\n")
	
	failCount := 0
	warnCount := 0
	okCount := 0
	
	for _, result := range results {
		switch result.Status {
		case checks.StatusOK:
			okCount++
		case checks.StatusWarn:
			warnCount++
		case checks.StatusFail:
			failCount++
		}
	}
	
	output.WriteString(fmt.Sprintf("**Summary:** %d passed ✅ | %d warnings ⚠️ | %d failures ❌\n\n", okCount, warnCount, failCount))
	
	output.WriteString("## Check Details\n\n")
	output.WriteString("| Check | Status | Message |\n")
	output.WriteString("|-------|--------|----------|\n")
	
	for _, result := range results {
		var icon string
		switch result.Status {
		case checks.StatusOK:
			icon = "✅"
		case checks.StatusWarn:
			icon = "⚠️"
		case checks.StatusFail:
			icon = "❌"
		case checks.StatusSkip:
			icon = "⏭️"
		}
		
		message := strings.ReplaceAll(result.Message, "\n", "<br>")
		output.WriteString(fmt.Sprintf("| %s | %s | %s |\n", result.Name, icon, message))
	}
	
	// Add detailed information for failures and warnings
	hasDetails := false
	for _, result := range results {
		if result.Details != "" && (result.Status == checks.StatusFail || result.Status == checks.StatusWarn) {
			if !hasDetails {
				output.WriteString("\n## Detailed Information\n\n")
				hasDetails = true
			}
			output.WriteString(fmt.Sprintf("### %s\n\n", result.Name))
			output.WriteString(fmt.Sprintf("**Status:** %s\n\n", result.Status))
			output.WriteString(fmt.Sprintf("**Message:** %s\n\n", result.Message))
			output.WriteString(fmt.Sprintf("**Details:**\n```\n%s\n```\n\n", result.Details))
		}
	}
	
	if failCount > 0 || warnCount > 0 {
		output.WriteString("## Remediation\n\n")
		output.WriteString("Run the following command locally to fix issues:\n")
		output.WriteString("```bash\n")
		output.WriteString("devpreflight fix --env\n")
		output.WriteString("```\n")
	}
	
	return output.String()
}
