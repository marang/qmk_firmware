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
- go/feature-renode: added Renode platform and monitor scripts for STM32L432 keyboard with USB HID.

- work: added Makefile and Renode launch script with CI logging.
- work: added CI workflow to build TinyGo firmware, run Renode tests, and archive HID logs.
- work: added STM32 DFU bootloader package, flash rule, and recovery docs.

- work: added power package with sleep loop and tests.

- go/feature-keymap: added keymap package with parser and macros.
- work: added debug logging package and instrumented drivers.
- work: added README, CONTRIBUTING, and API docs.
- go/feature-stm32l432-pin: implemented register-level pin, SPI, I2C, USB and tests.
- go/feature-keychron-pinconfig: defined GPIO constants and configured peripheral pins for Keychron Q1 Pro.
- go/feature-hid-usage: populated keycode lookup with full HID usage IDs and added tests.
- work: validated RGB LED count parameter and tests for invalid inputs.
- work: verified Renode command existence and allowed path override via RENODE.
- work: ensured run-renode script checks for renode command and supports RENODE override.
- work: updated TinyGo CI to use acifani/setup-tinygo@v2 and actions/upload-artifact@v4.
- go/feature-ci-act: documented Docker/act usage and added run-act script verifying artifacts.
- work: updated docs workflow to use qmk docs build command.
- go/feature-renode-install: added script to install Renode portable build and updated CI workflow.
- work: replaced qmk docs --build with qmk generate docs --build and ensured qmk CLI update in docs workflow.
- work: configured labeler workflow permissions and label syncing.
- work: added scripts label rule and inlined labeler workflow permissions.
- work: replaced spaces with tabs in Makefile recipes and verified build with Renode.
- work: set Makefile tinygo target to local board file and added requirements-dev.txt.

- work: prefixed GitHub Action workflow names with go/main label.
- work: added Keychron Q1 Pro TinyGo target file.
- work: ensured run-renode script exports GOROOT and namespaced local device/runtime packages.
- work: defined STM32L432 peripheral pin constants to resolve TinyGo build errors.
- go/feature-uart: stubbed UART and USB serial interfaces for STM32L432.
- work: replaced TinyGo machine package to expose STM32L432 peripheral pin constants.
- work: namespaced machine module to avoid import ambiguity with TinyGo runtime.
- work: installed TinyGo 0.39.0 and make build fails: undefined peripheral pin constants in std machine package.
- work: added minimal machine stubs and Makefile copy step to fix TinyGo build.
- work: staged TinyGo root locally during build to avoid system writes.
