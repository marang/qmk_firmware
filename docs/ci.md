# Continuous Integration

This guide explains how to run the TinyGo/Renode CI workflow locally.

## Install Docker

Docker is required for running GitHub Actions locally. On Debian/Ubuntu:

```sh
sudo apt-get update
sudo apt-get install --yes docker.io
sudo systemctl start docker
sudo usermod -aG docker "$USER"
# Log out and back in so group changes take effect
```

Verify that Docker is available:

```sh
docker info
```

## Install `act`

[`act`](https://github.com/nektos/act) executes GitHub Actions workflows locally.
Install it using your package manager or the project's install script. On Debian/Ubuntu:

```sh
curl -L https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash
```

## Run the workflow

From the repository root, execute:

```sh
act -W .github/workflows/tinygo-renode.yml -j build
```

This runs the `build` job defined in `.github/workflows/tinygo-renode.yml`.
The `scripts/run-act.sh` helper script checks that the Docker daemon is
running, executes the workflow with a suitable image, and verifies that
`hid.log` and `build/firmware.elf` are produced.
