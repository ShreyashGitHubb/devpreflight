package checks

import (
	"fmt"
	"os"
	"strings"

	"github.com/devpreflight/devpreflight/internal/config"
	"github.com/joho/godotenv"
)

type EnvParityChecker struct{}

func (c *EnvParityChecker) Name() string {
	return "env_parity"
}

func (c *EnvParityChecker) Enabled(cfg *config.Config) bool {
	return cfg.Checks.EnvParity
}

func (c *EnvParityChecker) Run(cfg *config.Config) CheckResult {
	// Check if .env.example exists
	if _, err := os.Stat(".env.example"); os.IsNotExist(err) {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusWarn,
			Message: ".env.example not found",
			Details: "Consider creating .env.example to document required environment variables",
		}
	}
	
	// Load .env.example
	exampleEnv, err := godotenv.Read(".env.example")
	if err != nil {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusFail,
			Message: "Failed to parse .env.example",
			Details: err.Error(),
		}
	}
	
	// Check if .env exists
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusFail,
			Message: ".env file missing",
			Details: fmt.Sprintf("Create .env with %d required keys from .env.example", len(exampleEnv)),
		}
	}
	
	// Load .env
	actualEnv, err := godotenv.Read(".env")
	if err != nil {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusFail,
			Message: "Failed to parse .env",
			Details: err.Error(),
		}
	}
	
	// Find missing and extra keys
	var missing []string
	var suspicious []string
	
	for key := range exampleEnv {
		if _, exists := actualEnv[key]; !exists {
			missing = append(missing, key)
		}
	}
	
	// Check for suspicious values (potential secrets in plain text)
	for key, value := range actualEnv {
		if len(value) > 50 && (strings.Contains(value, "==") || isLongHex(value)) {
			suspicious = append(suspicious, key)
		}
	}
	
	if len(missing) > 0 {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusFail,
			Message: fmt.Sprintf("%d missing environment variables", len(missing)),
			Details: fmt.Sprintf("Missing keys: %s\nRun 'devpreflight fix --env' to add placeholders", strings.Join(missing, ", ")),
		}
	}
	
	if len(suspicious) > 0 {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusWarn,
			Message: "Potential secrets detected",
			Details: fmt.Sprintf("Keys with long encoded values: %s", strings.Join(suspicious, ", ")),
		}
	}
	
	return CheckResult{
		Name:    c.Name(),
		Status:  StatusOK,
		Message: ".env matches .env.example",
		Details: fmt.Sprintf("%d environment variables validated", len(exampleEnv)),
	}
}

func isLongHex(s string) bool {
	if len(s) < 32 {
		return false
	}
	for _, r := range s {
		if !((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')) {
			return false
		}
	}
	return true
}
