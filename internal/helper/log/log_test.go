package log

import "testing"

// TestConsoleLog console log.
func TestConsoleLog(t *testing.T) {
	c := new(Console)

	c.Debug("Hello")
	c.Info("info")
	c.Warn("warn")
	c.Error("error")
}
