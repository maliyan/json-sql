// Package log console 形式日志输出.
package log

import (
	"github.com/fatih/color"
)

// ILog log interface.
type ILog interface {
	// Debug debug log function
	Debug(input string)

	// Info info log.
	Info(input string)

	// Warn warning log.
	Warn(input string)

	// Error error log.
	Error(input string)
}

// Console console output ILog.
type Console struct{}

// ConsoleDebug invoke new Console.Debug.
func ConcoleDebug(input string) {
	new(Console).Debug(input)
}

// ConsoleInfo invoke Console.Info.
func ConsoleInfo(input string) {
	new(Console).Info(input)
}

// ConsoleWarn invoke Console.Warn.
func ConsoleWarn(input string) {
	new(Console).Warn(input)
}

// ConsoleErr invoke Console.Error.
func ConsoleErr(input string) {
	new(Console).Error(input)
}

// Debug console debug log.
func (c *Console) Debug(input string) {
	color.White("Debug: %v", input)
}

// Info console info log.
func (c *Console) Info(input string) {
	color.Cyan("Info: %v", input)
}

// Warn console warn log.
func (c *Console) Warn(input string) {
	color.Yellow("Warn: %v", input)
}

// Error console error log.
func (c *Console) Error(input string) {
	color.Red("Err: %v", input)
}
