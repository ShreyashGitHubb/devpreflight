# devpreflight upgrade

The `upgrade` command allows you to update your local DevPreflight installation to the latest version available on GitHub.

## Usage

```bash
devpreflight upgrade
```

## Description

This command will:
1. Check for a new version on GitHub
2. Download the appropriate binary for your platform
3. Replace the current binary with the new version
4. Preserve executable permissions

## Examples

```bash
# Update to the latest version
$ devpreflight upgrade
Upgrading from v0.1.0 to v0.2.0...
Successfully upgraded to v0.2.0!

# When already on the latest version
$ devpreflight upgrade
You're already running the latest version!
```