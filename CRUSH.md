# Guidance for future agents. Read before working and append concise notes to
**Task Notes** after each run.

## Use `CRUSH-dep.md` to record missing packages, dependencies, or process improvements encountered during tasks.

## Understand the task and decide wise to be efficient
- Distinguish between updating documentation and code
- Code needs to run tests and other tasks
- Updating a documentation does not need to run tests or compiling or similar executions.

## Build and Lint
- Use the Go toolchain for building and testing.
- Format code with `go fmt ./...`.
- Run `go test ./...` and include the output in the final report.


## Git Workflow
- Create a feature branch with following naming convetion `go/feature-*` for changes and open a PR for review; merge to `go/main`
  only after approval.
- The development is tracked in `go/*`; prefix related branches with `go/feature-*`.
- Keep commits focused and ensure the working tree is clean before committing.

## Maintenance
- Keep this document brief, adding only noteworthy guidance.
- Add nested `CRUSH.md` files for new directories when useful.
- Record missing packages, dependencies, or process improvements here to avoid repeated issues.

## Task Notes
When adding a task note, begin with the current branch name. For example: `go/feature-*: describe task`
- Initial guidelines added.
- Clarified branch workflow and emphasis on concise updates.
- Submodules updated for Q1 Pro work.
- go/feature-stm32l432: added STM32L432 TinyGo support and board docs.

- work: added TinyGo matrix scanner using 74HC595 shift register and logging.
- work: integrated HID package mapping matrix events to USB reports.

- work: added RGB LED strip with brightness control and tests.
