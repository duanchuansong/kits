package xlog

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Info("Testing Info function")

	// Add assertions here if needed
}

func TestInfof(t *testing.T) {
	Infof("Testing Infof function with template: %s", "test")
	// Add assertions here if needed
}

func TestInfow(t *testing.T) {
	Infow("Testing Infow function with key-value pairs")
}

func TestWarn(t *testing.T) {
	Warn("Testing Warn function")
	// Add assertions here if needed
}

func TestWarnf(t *testing.T) {
	Infof("Testing Warnf function with template: %s", "test")
	// Add assertions here if needed
}

func TestDebug(t *testing.T) {
	Debug("Testing Debug function")
	// Add assertions here if needed
}

func TestDebugf(t *testing.T) {
	Debugf("Testing Debugf function with template: %s", "test")
	// Add assertions here if needed
}

func TestError(t *testing.T) {
	Error("Testing Error function")
	// Add assertions here if needed
}

func TestErrorf(t *testing.T) {
	Errorf("Testing Errorf function with template: %s", "test")
	// Add assertions here if needed
}
