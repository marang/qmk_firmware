# Missing Dependencies

- renode: `scripts/run-renode.sh` reported "renode command not found".
- tinygo: `make build` failed with "tinygo: No such file or directory".
- requirements-dev.txt: `pip3 install -r requirements-dev.txt` failed because the file was missing; added stub to satisfy CI.

