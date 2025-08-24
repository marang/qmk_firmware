# Renode Installation

Renode can be installed either automatically using the provided script or manually by downloading a portable build.

## Using the install script

Run the helper script from the repository root:

```bash
scripts/install-renode.sh
```

The script downloads the latest portable build from [builds.renode.io](https://builds.renode.io/), extracts it to `~/renode`, and makes the `renode` binary available on your `PATH`. When used in GitHub Actions, the script updates `GITHUB_PATH` so subsequent steps can invoke `renode`.

## Manual installation

1. Download the latest portable build:
   ```bash
   curl -LO https://builds.renode.io/renode-latest.linux-portable.tar.gz
   ```
2. Extract the archive to a directory:
   ```bash
   mkdir -p $HOME/renode
   tar -xzf renode-latest.linux-portable.tar.gz -C $HOME/renode --strip-components=1
   ```
3. Add the directory to your `PATH`:
   ```bash
   export PATH="$HOME/renode:$PATH"
   ```
4. Verify the installation:
   ```bash
   renode --version
   ```

This will make the `renode` command available for running simulations.
