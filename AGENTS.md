# AGENTS Instructions

This file collects guidance for future agents working in this repository.
Always read it before making changes, and append any new information or
lessons learned after completing tasks.

## Build and Lint
- Use the `qmk` CLI for all QMK tasks.
- For keyboard-related changes, run:
  - `qmk lint -kb <keyboard>`
  - `qmk compile -kb <keyboard> -km <keymap>`
- Include the output of these commands in the final report.

## Git Workflow
- Work directly on the `master` branch; avoid creating new branches.
- Keep commits focused and ensure the working tree is clean before committing.

## Maintenance
- If new conventions or requirements arise, update this `AGENTS.md`
  with concise notes so future runs benefit from the knowledge.
- When adding new directories, consider whether a nested `AGENTS.md`
  would help clarify instructions specific to that scope.

## Task Notes
- Initial guidelines added.
