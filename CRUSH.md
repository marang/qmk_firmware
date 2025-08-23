# Guidance for future agents. Read before working and append concise notes to
**Task Notes** after each run.

## Build and Lint
- `qmk` CLI is not installed by default; install it with `pip install qmk` before running `qmk` commands.
- Use the `qmk` CLI for all QMK tasks.
- For keyboard-related changes, run:
  - `qmk lint -kb <keyboard>`
  - `qmk compile -kb <keyboard> -km <keymap>`
- Include the output of these commands in the final report.

## Git Workflow
- Create a feature branch for changes and open a PR for review; merge to `master`
  only after approval.
- Keep commits focused and ensure the working tree is clean before committing.

## Maintenance
- Keep this document brief, adding only noteworthy guidance.
- Add nested `CRUSH.md` files for new directories when useful.
- Record missing packages, dependencies, or process improvements here to avoid repeated issues.

## Go Firmware Branch
- Go firmware development is tracked in `go/feature-q1pro`; prefix related branches with `go/`.

## Task Notes
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

