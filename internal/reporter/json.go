package reporter

import (
	"encoding/json"
	"fmt"

	"github.com/devpreflight/devpreflight/internal/checks"
)

type JSONReporter struct{}

func NewJSONReporter() *JSONReporter {
	return &JSONReporter{}
}

type JSONReport struct {
	Summary  string               `json:"summary"`
	Checks   []checks.CheckResult `json:"checks"`
	ExitCode int                  `json:"exit_code"`
}

func (r *JSONReporter) Report(results []checks.CheckResult) string {
	failCount := 0
	warnCount := 0
	
	for _, result := range results {
		if result.Status == checks.StatusFail {
			failCount++
		} else if result.Status == checks.StatusWarn {
			warnCount++
		}
	}
	
	exitCode := 0
	if failCount > 0 {
		exitCode = 2
	} else if warnCount > 0 {
		exitCode = 1
	}
	
	summary := fmt.Sprintf("%d warnings, %d failures", warnCount, failCount)
	
	report := JSONReport{
		Summary:  summary,
		Checks:   results,
		ExitCode: exitCode,
	}
	
	data, _ := json.MarshalIndent(report, "", "  ")
	return string(data)
}
