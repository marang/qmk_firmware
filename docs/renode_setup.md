# Renode Setup

Renode version 1.15 or later is required for virtual keyboard testing in QMK.

## Linux

```bash
sudo apt install renode
```

Other distributions may provide a `renode` package or you can download the latest release from [renode/renode](https://github.com/renode/renode/releases).

## macOS

```bash
brew install --cask renode
```

## Windows

Use [Chocolatey](https://chocolatey.org/) to install:

```powershell
choco install renode
```

## Verify Installation

After installation, check the version:

```bash
renode --version
```

Ensure the reported version is 1.15 or newer.

For other platforms or installation options, consult the [official Renode documentation](https://renode.io/).
