# DevPreflight CLI

## Overview

DevPreflight is a preflight CLI tool designed to run curated checks before deployment. It validates environment parity, performs Dockerfile linting, validates Kubernetes manifests, checks observability hooks, and detects flaky tests. The tool is distributed as a single static binary and produces human-readable reports for developers and CI/CD pipelines.

**Current Status:** MVP Complete (October 2025)
- All core checks implemented and tested
- Multiple output formats supported (console, JSON, Markdown)
- Auto-fix functionality for environment variable parity
- Example fixtures and working demo available

## User Preferences

Preferred communication style: Simple, everyday language.

## System Architecture

### Core Technology Stack

**Language & Distribution:**
- **Primary Language**: Go
  - Rationale: Enables creation of static binaries for easy distribution, strong CLI tooling ecosystem, and straightforward cross-platform compilation
  - Alternative considered: Node.js (rejected due to worse CI install UX)
- **Binary Distribution**: Single static executable (`devpreflight`)

**CLI Framework:**
- **Framework**: Cobra
  - Purpose: Provides structured command handling, subcommand routing, and flag parsing
  - Subcommands: `check`, `fix`, `ci-report`

**Configuration Management:**
- **Config File**: `.devpreflightrc.yml`
- **Config Library**: Viper
  - Purpose: Handles configuration file parsing and runtime overrides
  - Enables hierarchical configuration (file + flags + environment variables)

### Application Design Patterns

**Wrapper-Based Architecture:**
- Primary approach: Wrap existing, proven linters and scanners rather than reimplementing validation logic
- External tools integrated:
  - `hadolint` - Dockerfile linting
  - `kubeval`/`kubeconform` - Kubernetes manifest validation
  - `trivy` - Security scanning
- Implementation: Shell out to external binaries (acceptable for MVP to reduce development time)

**Check Categories:**
1. Environment parity validation
2. Dockerfile quick-lints
3. Kubernetes manifest validation
4. Observability hooks verification
5. Flaky test detection

**Output Formats:**
- Human-readable reports for developer consumption
- CI/CD-friendly formatted output for pipeline integration

## External Dependencies

### Development Dependencies

**Go Modules:**
- `github.com/spf13/cobra@latest` - CLI framework and command structure
- `github.com/spf13/viper@latest` - Configuration management

### Runtime Dependencies (External Binaries)

**Linting & Validation Tools:**
- **hadolint** - Dockerfile best practices and linting
- **kubeval** or **kubeconform** - Kubernetes YAML validation
- **trivy** - Container and security scanning

**Integration Method:**
- External binaries invoked via shell execution
- Tools are expected to be available in system PATH or bundled with distribution
- Exit codes and stdout/stderr captured for report generation

**Note**: The wrapper approach allows rapid MVP development while maintaining flexibility to replace or internalize specific checks in future iterations.