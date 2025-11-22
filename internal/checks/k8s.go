package checks

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/devpreflight/devpreflight/internal/config"
	"go.yaml.in/yaml/v3"
)

type K8sChecker struct{}

// NewK8sChecker creates a new K8sChecker instance
func NewK8sChecker() *K8sChecker {
	return &K8sChecker{}
}

func (c *K8sChecker) Name() string {
	return "k8s_validate"
}

func (c *K8sChecker) Enabled(cfg *config.Config) bool {
	return cfg.Checks.K8sValidate
}

func (c *K8sChecker) Run(cfg *config.Config) CheckResult {
	// Look for k8s manifests in common locations
	patterns := []string{
		"*.yaml",
		"*.yml",
		"k8s/*.yaml",
		"k8s/*.yml",
		"kubernetes/*.yaml",
		"kubernetes/*.yml",
		"deploy/*.yaml",
		"deploy/*.yml",
	}
	
	var manifestFiles []string
	for _, pattern := range patterns {
		matches, _ := filepath.Glob(pattern)
		manifestFiles = append(manifestFiles, matches...)
	}
	
	if len(manifestFiles) == 0 {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusSkip,
			Message: "No Kubernetes manifests found",
			Details: "Skipping K8s validation",
		}
	}
	
	var issues []string
	var warnings []string
	validManifests := 0
	
	for _, file := range manifestFiles {
		data, err := os.ReadFile(file)
		if err != nil {
			continue
		}
		
		// Try to parse as K8s manifest
		var manifest map[string]interface{}
		if err := yaml.Unmarshal(data, &manifest); err != nil {
			continue
		}
		
		// Check if it looks like a K8s resource
		apiVersion, hasAPI := manifest["apiVersion"].(string)
		kind, hasKind := manifest["kind"].(string)
		
		if !hasAPI || !hasKind {
			continue
		}
		
		validManifests++
		
		// Check for deprecated API versions
		deprecated := []string{
			"apps/v1beta1",
			"apps/v1beta2",
			"extensions/v1beta1",
			"networking.k8s.io/v1beta1",
		}
		
		for _, dep := range deprecated {
			if strings.Contains(apiVersion, dep) {
				issues = append(issues, fmt.Sprintf("%s: Deprecated apiVersion '%s' (kind: %s)", file, apiVersion, kind))
			}
		}
		
		// Check for missing metadata.name
		if metadata, ok := manifest["metadata"].(map[string]interface{}); ok {
			if _, hasName := metadata["name"]; !hasName {
				issues = append(issues, fmt.Sprintf("%s: Missing metadata.name", file))
			}
		} else {
			issues = append(issues, fmt.Sprintf("%s: Missing metadata section", file))
		}
		
		// Deployment-specific checks
		if kind == "Deployment" {
			if spec, ok := manifest["spec"].(map[string]interface{}); ok {
				if _, hasSelector := spec["selector"]; !hasSelector {
					issues = append(issues, fmt.Sprintf("%s: Deployment missing spec.selector", file))
				}
			}
		}
		
		// Check for hostPath usage (warning)
		if spec, ok := manifest["spec"].(map[string]interface{}); ok {
			specStr := fmt.Sprintf("%v", spec)
			if strings.Contains(specStr, "hostPath") {
				warnings = append(warnings, fmt.Sprintf("%s: Using hostPath volume (security concern)", file))
			}
		}
	}
	
	if validManifests == 0 {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusSkip,
			Message: "No valid Kubernetes manifests found",
			Details: fmt.Sprintf("Checked %d YAML files", len(manifestFiles)),
		}
	}
	
	if len(issues) > 0 {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusFail,
			Message: fmt.Sprintf("%d critical issues in %d manifests", len(issues), validManifests),
			Details: strings.Join(issues, "\n"),
		}
	}
	
	if len(warnings) > 0 {
		return CheckResult{
			Name:    c.Name(),
			Status:  StatusWarn,
			Message: fmt.Sprintf("%d warnings in %d manifests", len(warnings), validManifests),
			Details: strings.Join(warnings, "\n"),
		}
	}
	
	return CheckResult{
		Name:    c.Name(),
		Status:  StatusOK,
		Message: fmt.Sprintf("Validated %d Kubernetes manifests", validManifests),
		Details: "All manifests look good",
	}
}
