# 🚀 DevPreflight

DevPreflight is a powerful CLI tool that helps you catch configuration issues and enforce best practices in your development workflow before they cause problems in production.

## 🌟 Features

- ✅ Environment variable parity checking
- 🐳 Dockerfile best practices validation
- ⚓ Kubernetes manifest validation
- 🔍 Flaky test detection
- 📊 Observability checks
- 🛠️ Auto-fix capabilities for common issues

## 🚀 Quick Start

```bash
# Install using Go
go install github.com/devpreflight/devpreflight@latest

# Run checks in your project
devpreflight check

# Generate documentation
devpreflight docs

# View changelog
devpreflight changelog
```

## 📘 Commands

- `check` - Run all preflight checks
- `fix` - Auto-fix detected issues
- `ci-report` - Generate CI-friendly reports
- `version` - Print version information
- `docs` - Generate Markdown documentation
- `man` - Generate Unix man pages
- `changelog` - View the changelog
- `completion` - Generate shell completion scripts

## 🛠️ Installation

### Using Go

```bash
go install github.com/devpreflight/devpreflight@latest
```

### Binary Releases

Download the latest binary for your platform from our [releases page](https://github.com/devpreflight/devpreflight/releases).

## 🧪 Example Usage

Here's an example of running checks on a project:

```bash
$ devpreflight check
✓ Environment Variables: All required variables present
✓ Dockerfile: Best practices validated
⚠ Kubernetes Manifests: Resource limits not set
✗ Tests: 2 flaky tests detected
```

## 📚 Documentation

Full documentation is available at [https://devpreflight.github.io/devpreflight](https://devpreflight.github.io/devpreflight)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.