# Changelog

All notable changes to DevPreflight will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2025-10-07

### Added
- Initial release of DevPreflight CLI
- Environment parity checker (.env vs .env.example validation)
- Dockerfile linting (checks for :latest tags, multi-stage builds)
- Kubernetes manifest validation (deprecated APIs, schema validation)
- Observability instrumentation detection (OTEL, Sentry, Prometheus)
- Flaky test detector (runs tests multiple times)
- Multiple output formats: console (colorized), JSON, Markdown
- Auto-fix functionality for environment variable parity
- Configuration file support (.devpreflightrc.yml)
- Exit codes: 0 (success), 1 (warnings), 2 (failures)
- Comprehensive help documentation for all commands
- Shell completion support (bash, zsh, fish, powershell)
- Automatic documentation generation (Markdown and man pages)
- Verbose mode for detailed output
- CI/CD integration examples

### Security
- Never auto-populates secrets or credentials
- Uses __REPLACE_ME__ placeholders for missing environment variables
- Safe auto-fix operations only

## [Unreleased]

### Planned Features
- Plugin system for custom checks
- Additional auto-fix capabilities for Dockerfile issues
- Enhanced observability detection (source code scanning)
- Test coverage analysis
- Security vulnerability scanning integration
- Docker image size optimization suggestions
- Kubernetes best practices enforcement
