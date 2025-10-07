# DevPreflight Documentation

Welcome to the DevPreflight documentation! This tool helps you catch configuration issues and enforce best practices in your development workflow.

## Core Features

### 🔍 Environment Variable Parity
Ensures consistency between different environment configurations (development, staging, production).
[Learn more](devpreflight_check.md#environment-parity)

### 🐳 Dockerfile Validation
Validates Dockerfile best practices and common pitfalls.
[Learn more](devpreflight_check.md#dockerfile-validation)

### ⚓ Kubernetes Manifest Validation
Checks Kubernetes manifests for best practices and common configuration issues.
[Learn more](devpreflight_check.md#kubernetes-validation)

### 🧪 Flaky Test Detection
Identifies potentially flaky tests in your test suite.
[Learn more](devpreflight_check.md#test-validation)

### 📊 Observability Checks
Validates logging, metrics, and tracing configurations.
[Learn more](devpreflight_check.md#observability)

## Commands

- [check](devpreflight_check.md) - Run all preflight checks
- [fix](devpreflight_fix.md) - Auto-fix detected issues
- [ci-report](devpreflight_ci-report.md) - Generate CI-friendly reports
- [upgrade](devpreflight_upgrade.md) - Upgrade to the latest version
- [docs](devpreflight_docs.md) - Generate documentation
- [man](devpreflight_man.md) - Generate man pages
- [version](devpreflight_version.md) - Print version information
- [changelog](devpreflight_changelog.md) - View the changelog

## Installation

### Using Go

```bash
go install github.com/devpreflight/devpreflight@latest
```

### Using Homebrew

```bash
brew tap devpreflight/devpreflight
brew install devpreflight
```

### Manual Installation

Download the latest binary for your platform from our [releases page](https://github.com/devpreflight/devpreflight/releases).

## Configuration

DevPreflight can be configured using a `devpreflight.yaml` file in your project root. Here's a sample configuration:

```yaml
checks:
  env_parity:
    enabled: true
    environments:
      - dev
      - staging
      - prod
  dockerfile:
    enabled: true
  kubernetes:
    enabled: true
  flaky_tests:
    enabled: true
    threshold: 0.1
  observability:
    enabled: true
```

## Examples

Check out our [examples directory](https://github.com/devpreflight/devpreflight/tree/main/examples) for sample configurations and use cases.