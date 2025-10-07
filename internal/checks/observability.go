package checks

import (
        "fmt"
        "os"
        "strings"

        "github.com/devpreflight/devpreflight/internal/config"
)

type ObservabilityChecker struct{}

func (c *ObservabilityChecker) Name() string {
        return "observability"
}

func (c *ObservabilityChecker) Enabled(cfg *config.Config) bool {
        return cfg.Checks.Observability
}

func (c *ObservabilityChecker) Run(cfg *config.Config) CheckResult {
        // Check for common observability environment variables
        otelVars := []string{
                "OTEL_EXPORTER_OTLP_ENDPOINT",
                "OTEL_TRACES_SAMPLER",
                "OTEL_SERVICE_NAME",
        }
        
        sentryVars := []string{
                "SENTRY_DSN",
        }
        
        prometheusVars := []string{
                "PROMETHEUS_MULTIPROC_DIR",
        }
        
        var foundOTEL, foundSentry, foundPrometheus bool
        
        // Check OTEL
        for _, v := range otelVars {
                if os.Getenv(v) != "" {
                        foundOTEL = true
                        break
                }
        }
        
        // Check Sentry
        for _, v := range sentryVars {
                if os.Getenv(v) != "" {
                        foundSentry = true
                        break
                }
        }
        
        // Check Prometheus
        for _, v := range prometheusVars {
                if os.Getenv(v) != "" {
                        foundPrometheus = true
                        break
                }
        }
        
        // Also check .env.example for these variables
        if data, err := os.ReadFile(".env.example"); err == nil {
                content := string(data)
                if !foundOTEL && containsAny(content, otelVars) {
                        foundOTEL = true
                }
                if !foundSentry && containsAny(content, sentryVars) {
                        foundSentry = true
                }
                if !foundPrometheus && containsAny(content, prometheusVars) {
                        foundPrometheus = true
                }
        }
        
        var found []string
        if foundOTEL {
                found = append(found, "OpenTelemetry")
        }
        if foundSentry {
                found = append(found, "Sentry")
        }
        if foundPrometheus {
                found = append(found, "Prometheus")
        }
        
        if len(found) == 0 {
                return CheckResult{
                        Name:    c.Name(),
                        Status:  StatusWarn,
                        Message: "No observability instrumentation detected",
                        Details: "Consider adding OpenTelemetry, Sentry, or Prometheus metrics.\n" +
                                "Example OTEL env vars:\n" +
                                "  OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318\n" +
                                "  OTEL_SERVICE_NAME=my-service\n" +
                                "  OTEL_TRACES_SAMPLER=always_on",
                }
        }
        
        if len(found) < 2 {
                return CheckResult{
                        Name:    c.Name(),
                        Status:  StatusWarn,
                        Message: fmt.Sprintf("Limited observability: only %s configured", strings.Join(found, ", ")),
                        Details: "Consider adding more observability tools for better production monitoring",
                }
        }
        
        return CheckResult{
                Name:    c.Name(),
                Status:  StatusOK,
                Message: fmt.Sprintf("Observability configured: %s", strings.Join(found, ", ")),
                Details: fmt.Sprintf("%d observability systems detected", len(found)),
        }
}

func containsAny(content string, vars []string) bool {
        for _, v := range vars {
                if strings.Contains(content, v) {
                        return true
                }
        }
        return false
}
