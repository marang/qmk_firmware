# Guidance for future agents. Read before working and append concise notes to
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
- Add nested `CRUSH.md` files for new directories when useful.

## Task Notes
- Initial guidelines added.
- Clarified branch workflow and emphasis on concise updates.
- Submodules updated for Q1 Pro work.

# QMK Keyboard Development Guidelines

## Build Commands
- `make` - Build all keyboards
- `make <keyboard>:<keymap>` - Build a specific keyboard/keymap
- `make <keyboard>:<keymap>:flash` - Build and flash a keyboard
- `make <keyboard>:<keymap>:clean` - Clean build files for a keyboard
- `qmk compile -kb <keyboard> -km <keymap>` - QMK CLI compile command
- `qmk new-keymap -kb <keyboard> -km <keymap>` - Create a new keymap

## Test Commands
- `make test` - Run all tests
- `make test:<test_name>` - Run a specific test
- `make test:<test_name>:clean` - Clean a specific test
- `make test:all` - Run all tests

## Lint Commands
- `make lint` - Run linting checks
- `make format` - Format code according to style guidelines

## Code Style Guidelines
- Follow QMK's existing code style and conventions
- Use C89 style for embedded C programming
- All code must compile with GCC
- Use descriptive variable and function names
- Follow existing naming conventions in the codebase
- Maintain consistency with existing code structure
- Use proper commenting for complex logic
- All code must be compatible with the QMK build system

## Git Submodules
- Always run `git submodule update --init --recursive` after cloning or pulling changes
- This ensures all required dependencies (ChibiOS, etc.) are properly initialized

## Common Pitfalls and Solutions

### 1. API Compatibility Issues
**Problem**: Legacy QMK code using deprecated functions that don't exist in modern QMK versions.
**Examples**:
- `rgb_matrix_disable_timeout_set()` → Function doesn't exist
- `rgb_matrix_disable_time_reset()` → Function doesn't exist
- `rgb_matrix_is_driver_shutdown()` → Function doesn't exist
- `rgb_matrix_driver_allow_shutdown()` → Function doesn't exist

**Solution**:
- Check current QMK API documentation
- Use conditional compilation with `#ifdef` guards
- Comment out or replace deprecated functions with modern equivalents
- When in doubt, simplify functionality rather than trying to reimplement missing APIs

### 2. Header File Issues
**Problem**: Missing or incorrect header includes causing compilation errors.
**Example**: `#include "nvm.h"` → File doesn't exist

**Solution**:
- Check the actual file structure in `quantum/` directory
- Use correct paths like `#include "nvm/eeprom/nvm_eeprom_eeconfig_internal.h"`
- Verify header file existence before including

### 3. Configuration Constant Issues
**Problem**: Constants like `EECONFIG_RGB_MATRIX` or `RGB_MATRIX_TIMEOUT_INFINITE` not defined.
**Solution**:
- Check if constants exist in current QMK version
- Define fallback values when needed:
  ```c
  #ifndef RGB_MATRIX_TIMEOUT_INFINITE
  #    define RGB_MATRIX_TIMEOUT_INFINITE 0xFFFFFFFF
  #endif
  ```

### 4. Driver API Changes
**Problem**: RGB/LED matrix driver structure changed in modern QMK.
**Old Structure**:
```c
typedef struct {
    void (*init)(void);
    void (*set_color)(int index, uint8_t r, uint8_t g, uint8_t b);
    void (*set_color_all)(uint8_t r, uint8_t g, uint8_t b);
    void (*flush)(void);
    void (*exit_shutdown)(void);  // Removed in modern QMK
    void (*shutdown)(void);       // Removed in modern QMK
} rgb_matrix_driver_t;
```

**Modern Structure**:
```c
typedef struct {
    void (*init)(void);
    void (*set_color)(int index, uint8_t r, uint8_t g, uint8_t b);
    void (*set_color_all)(uint8_t r, uint8_t g, uint8_t b);
    void (*flush)(void);
    // shutdown functions removed
} rgb_matrix_driver_t;
```

**Solution**: 
- Remove calls to deprecated driver functions
- Use conditional compilation to handle API differences
- Simplify functionality when modern equivalents don't exist

### 5. Macro Definition Conflicts
**Problem**: Macros expanding to undefined functions causing compilation errors.
**Example**: `#define LED_DRIVER_ALLOW_SHUTDOWN rgb_matrix_driver_allow_shutdown` where function doesn't exist.

**Solution**:
- Comment out problematic macro definitions
- Use conditional compilation to check function existence
- Avoid defining macros that expand to non-existent functions

## Best Practices for Legacy Code Migration

### 1. Incremental Approach
- Fix one compilation error at a time
- Don't try to fix everything at once
- Test compilation frequently

### 2. Defensive Programming
- Use `#ifdef` guards for conditional compilation
- Provide fallback implementations for missing features
- Comment out problematic code rather than deleting it

### 3. Research Before Implementation
- Check current QMK documentation
- Look at working examples in the codebase
- Search for similar implementations in other keyboards

### 4. Focus on Core Functionality
- Get basic compilation working first
- Don't get bogged down in peripheral features
- Prioritize keyboard functionality over advanced features

## Troubleshooting Checklist

### When Compilation Fails:
1. **Check API compatibility** - Are you using deprecated functions?
2. **Verify header includes** - Do all headers exist and are correctly included?
3. **Check constant definitions** - Are all required constants defined?
4. **Review driver API** - Has the driver structure changed?
5. **Validate macro definitions** - Do macros expand to valid functions?

### When Functions Are Missing:
1. **Search QMK codebase** - Does the function exist with a different name?
2. **Check version compatibility** - Was the function removed in newer versions?
3. **Find alternatives** - Is there a modern equivalent?
4. **Implement workarounds** - Can you achieve the same result differently?

## Keyboard-Specific Notes

### Keychron Q1 Pro
- Uses STM32L432 processor
- Requires proper RGB matrix configuration
- Bluetooth functionality has complex dependencies
- Variants (ansi_knob, iso_knob, jis_encoder) have different configurations

## Testing
- Tests are located in the `tests/` directory
- Use the QMK test framework for unit testing
- All new features should include appropriate tests
- Test coverage should be maintained for existing functionality