//go:build !debug
// +build !debug

package debug

// Printf is a no-op when debug logging is disabled.
func Printf(format string, a ...interface{}) {}
