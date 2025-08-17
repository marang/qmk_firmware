# AGENTS Instructions

Guidance for future agents. Read before working and append concise notes to
**Task Notes** after each run.

## Build and Lint
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
- Add nested `AGENTS.md` files for new directories when useful.

## Task Notes
- Initial guidelines added.
- Clarified branch workflow and emphasis on concise updates.
