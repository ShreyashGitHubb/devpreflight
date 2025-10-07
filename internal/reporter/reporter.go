package reporter

import "github.com/devpreflight/devpreflight/internal/checks"

type Reporter interface {
	Report(results []checks.CheckResult) string
}
