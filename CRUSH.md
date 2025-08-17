# QMK Firmware Development Guidelines

## Build Commands
- `make` - Build all keyboards
- `make <keyboard>:<keymap>` - Build a specific keyboard/keymap
- `make <keyboard>:<keymap>:flash` - Build and flash a keyboard
- `make <keyboard>:<keymap>:clean` - Clean build files for a keyboard

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

## Testing
- Tests are located in the `tests/` directory
- Use the QMK test framework for unit testing
- All new features should include appropriate tests
- Test coverage should be maintained for existing functionality

## Git Submodules
- Always run `git submodule update --init --recursive` after cloning or pulling changes
- This ensures all required dependencies (ChibiOS, etc.) are properly initialized