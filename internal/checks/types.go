package checks

import "github.com/devpreflight/devpreflight/internal/config"

type CheckStatus string

const (
	StatusOK   CheckStatus = "ok"
	StatusWarn CheckStatus = "warn"
	StatusFail CheckStatus = "fail"
	StatusSkip CheckStatus = "skip"
)

type CheckResult struct {
	Name    string      `json:"name"`
	Status  CheckStatus `json:"status"`
	Details string      `json:"details"`
	Message string      `json:"message"`
}

type Checker interface {
	Run(cfg *config.Config) CheckResult
	Name() string
	Enabled(cfg *config.Config) bool
}

func RunAllChecks(cfg *config.Config) []CheckResult {
	checkers := []Checker{
		&EnvParityChecker{},
		&DockerfileChecker{},
		&K8sChecker{},
		&ObservabilityChecker{},
		&FlakyTestChecker{},
	}
	
	var results []CheckResult
	for _, checker := range checkers {
		if !checker.Enabled(cfg) {
			results = append(results, CheckResult{
				Name:   checker.Name(),
				Status: StatusSkip,
				Details: "Check disabled in config",
			})
			continue
		}
		
		result := checker.Run(cfg)
		results = append(results, result)
	}
	
	return results
}
