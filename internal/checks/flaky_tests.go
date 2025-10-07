package checks

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/devpreflight/devpreflight/internal/config"
)

type FlakyTestChecker struct{}

func (c *FlakyTestChecker) Name() string {
	return "flaky_tests"
}

func (c *FlakyTestChecker) Enabled(cfg *config.Config) bool {
	return cfg.Checks.FlakyTests
}

func (c *FlakyTestChecker) Run(cfg *config.Config) CheckResult {
	// Detect test runner
	testCmd := detectTestRunner()
	if testCmd == "" {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusSkip,
			Message: "No test runner detected",
			Details: "Skipping flaky test detection (supported: go test, pytest, npm test)",
		}
	}
	
	// Run tests twice to detect flakiness
	runs := 2
	var results []bool
	
	for i := 0; i < runs; i++ {
		passed := runTests(testCmd)
		results = append(results, passed)
	}
	
	// Check for inconsistent results
	allSame := true
	first := results[0]
	for _, r := range results {
		if r != first {
			allSame = false
			break
		}
	}
	
	if !allSame {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusFail,
			Message: "Flaky tests detected",
			Details: fmt.Sprintf("Tests produced inconsistent results across %d runs.\n"+
				"This indicates non-deterministic behavior.\n"+
				"Run '%s' multiple times to reproduce.", runs, testCmd),
		}
	}
	
	if !first {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusFail,
			Message: "Tests consistently failing",
			Details: fmt.Sprintf("Tests failed in all %d runs", runs),
		}
	}
	
	return CheckResult{
		Name:    c.Name(),
		Status:  StatusOK,
		Message: "No flaky tests detected",
		Details: fmt.Sprintf("Tests passed consistently across %d runs", runs),
	}
}

func detectTestRunner() string {
	// Check for Go
	if _, err := os.Stat("go.mod"); err == nil {
		if _, err := exec.LookPath("go"); err == nil {
			return "go test ./..."
		}
	}
	
	// Check for Python
	if fileExists("pytest.ini") || fileExists("setup.py") || fileExists("requirements.txt") {
		if _, err := exec.LookPath("pytest"); err == nil {
			return "pytest -q"
		}
	}
	
	// Check for Node.js
	if _, err := os.Stat("package.json"); err == nil {
		if _, err := exec.LookPath("npm"); err == nil {
			return "npm test --silent"
		}
	}
	
	return ""
}

func runTests(cmd string) bool {
	parts := strings.Fields(cmd)
	command := exec.Command(parts[0], parts[1:]...)
	command.Stdout = nil
	command.Stderr = nil
	
	err := command.Run()
	return err == nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
