package logx

import (
	"bytes"
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	t.Run("Test log with coloring enabled", func(t *testing.T) {
		ColoringEnabled = true
		expected := "\x1b[31m\x1b[42mTest Message\x1b[0m\n"
		buf := new(bytes.Buffer)
		log.SetOutput(buf)

		Log("Test Message", FGRED, BGGREEN)

		if got := buf.String(); got != expected {
			t.Errorf("Expected: %v, Got: %v", expected, got)
		}
	})

	t.Run("Test log with coloring disabled", func(t *testing.T) {
		ColoringEnabled = false
		expected := "Test Message\n"
		buf := new(bytes.Buffer)
		log.SetOutput(buf)

		Log("Test Message", FGRED, BGGREEN)

		if got := buf.String(); got != expected {
			t.Errorf("Expected: %v, Got: %v", expected, got)
		}
	})
}

func TestLogf(t *testing.T) {
	expected := "\x1b[31m\x1b[42mFormatted Message: 42\x1b[0m\n"
	buf := new(bytes.Buffer)
	log.SetOutput(buf)

	Logf("Formatted Message: %d", FGRED, BGGREEN, 42)

	if got := buf.String(); got != expected {
		t.Errorf("Expected: %v, Got: %v", expected, got)
	}
}

func TestLogWithLevel(t *testing.T) {
	t.Run("Test log with INFO level", func(t *testing.T) {
		expected := "\x1b[44mLog message with INFO level\x1b[0m\n"
		buf := new(bytes.Buffer)
		log.SetOutput(buf)

		LogWithLevel("Log message with INFO level", "INFO")

		if got := buf.String(); got != expected {
			t.Errorf("Expected: %v, Got: %v", expected, got)
		}
	})

	t.Run("Test log with WARNING level", func(t *testing.T) {
		expected := "\x1b[43mLog message with WARNING level\x1b[0m\n"
		buf := new(bytes.Buffer)
		log.SetOutput(buf)

		LogWithLevel("Log message with WARNING level", "WARNING")

		if got := buf.String(); got != expected {
			t.Errorf("Expected: %v, Got: %v", expected, got)
		}
	})
}

func TestLogM(t *testing.T) {
	expected := "\x1b[31m\x1b[42mTest\x1b[0m \x1b[34m\x1b[43mMultiple\x1b[0m \x1b[32m\x1b[44mMessages\x1b[0m\n"
	buf := new(bytes.Buffer)
	log.SetOutput(buf)

	LogM([]string{"Test", "Multiple", "Messages"}, FGRED, BGGREEN, FGBLUE, BGYELLOW, FGGREEN, BGBLUE)

	if got := buf.String(); got != expected {
		t.Errorf("Expected: %v, Got: %v", expected, got)
	}
}

func TestLogWithTimestamp(t *testing.T) {
	expected := "2023-01-14 15:04:05 \x1b[31m\x1b[42mTimestamped Message\x1b[0m\n"
	buf := new(bytes.Buffer)
	log.SetOutput(buf)

	LogWithTimestamp("Timestamped Message", FGRED, BGGREEN)

	if got := buf.String(); got != expected {
		t.Errorf("Expected: %v, Got: %v", expected, got)
	}
}
