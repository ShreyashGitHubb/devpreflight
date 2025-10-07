package checks

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/devpreflight/devpreflight/internal/config"
)

type DockerfileChecker struct{}

func (c *DockerfileChecker) Name() string {
	return "dockerfile_lint"
}

func (c *DockerfileChecker) Enabled(cfg *config.Config) bool {
	return cfg.Checks.DockerfileLint
}

func (c *DockerfileChecker) Run(cfg *config.Config) CheckResult {
	// Check if Dockerfile exists
	if _, err := os.Stat("Dockerfile"); os.IsNotExist(err) {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusSkip,
			Message: "No Dockerfile found",
			Details: "Skipping Dockerfile checks",
		}
	}
	
	file, err := os.Open("Dockerfile")
	if err != nil {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusFail,
			Message: "Failed to read Dockerfile",
			Details: err.Error(),
		}
	}
	defer file.Close()
	
	var issues []string
	var warnings []string
	hasMultiStage := false
	lineNum := 0
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		
		// Skip comments and empty lines
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		
		// Check for :latest tag
		if strings.HasPrefix(line, "FROM ") {
			if strings.Contains(line, ":latest") && cfg.Docker.ForbidLatest {
				issues = append(issues, fmt.Sprintf("Line %d: Using ':latest' tag (pin to specific version)", lineNum))
			}
			if strings.Contains(line, " AS ") || strings.Contains(line, " as ") {
				hasMultiStage = true
			}
		}
		
		// Check for problematic RUN patterns
		if strings.HasPrefix(line, "RUN ") {
			if strings.Contains(line, "apt-get install") && !strings.Contains(line, "--no-install-recommends") {
				warnings = append(warnings, fmt.Sprintf("Line %d: Consider using --no-install-recommends with apt-get", lineNum))
			}
			if strings.Contains(line, "curl") && !strings.Contains(line, "&&") && !strings.Contains(line, ";") {
				warnings = append(warnings, fmt.Sprintf("Line %d: Consider chaining commands to reduce layers", lineNum))
			}
		}
	}
	
	if err := scanner.Err(); err != nil {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusFail,
			Message: "Error scanning Dockerfile",
			Details: err.Error(),
		}
	}
	
	// Check for multi-stage build recommendation
	if !hasMultiStage && lineNum > 10 {
		warnings = append(warnings, "Consider using multi-stage builds to reduce image size")
	}
	
	if len(issues) > 0 {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusFail,
			Message: fmt.Sprintf("%d critical issues found", len(issues)),
			Details: strings.Join(issues, "\n"),
		}
	}
	
	if len(warnings) > 0 {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusWarn,
			Message: fmt.Sprintf("%d recommendations", len(warnings)),
			Details: strings.Join(warnings, "\n"),
		}
	}
	
	return CheckResult{
		Name:    c.Name(),
		Status:  StatusOK,
		Message: "Dockerfile looks good",
		Details: fmt.Sprintf("Scanned %d lines, no issues found", lineNum),
	}
}
