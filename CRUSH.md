# Guidance for future agents. Read before working and append concise notes to
**Task Notes** after each run.

## Use `CRUSH-dep.md` to record missing packages, dependencies, or process improvements encountered during tasks.

## Understand the task and decide wise to be efficient
- Distinguish between updating documentation and code
- Code needs to run tests and other tasks
- Updating a documentation does not need to run tests or compiling or similar executions.

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
- Removed unsupported encoder_map feature from Q1 Pro and tidied layout macros.
- Enabled VIA on Q1 Pro; lint still failing due to keyboard.json location.
- Removed stray firmware directory from Keychron Q1 Pro and added ignore rule.
- Set wear-leveling backing size and renamed keyboard.json to info.json for Q1 Pro; dropped VIA keymap and keyboard-level VIA to satisfy lint.
- Fixed bootloader include order and defined bootloader for Q1 Pro variants to resolve linking errors.

- Document missing packages and workflow improvements here; `qmk` CLI isn't installed by default and can be added with `pip install qmk`.
- Noted Go firmware branch `go/feature-q1pro` and branch prefix convention.

