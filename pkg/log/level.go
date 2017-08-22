// Copyright © 2017 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

package log

import (
	"errors"
	"strings"
)

var ErrInvalidLevel = errors.New("Invalid log level")

// Level is the level of logging.
type Level int8

const (
	// invalid is an invalid log level
	invalid Level = iota

	// DebugLevel is the log level for debug messages, usually turned of in production.
	DebugLevel

	// InfoLevel is the log level for informational messages.
	InfoLevel

	// Warn is the log level for warnings.
	// Warnings are more important than info but do not need individual human review.
	WarnLevel

	// Error is the log level for high priority error messages.
	// When everything is running smoothly, an application should not be logging error level messages.
	ErrorLevel

	// Fatal the log level for unrecoverable errors.
	FatalLevel
)

// String implements fmt.Stringer.
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	default:
		return "invalid"
	}
}

// ParseLevel parses a string into a log level or returns an error otherwise.
func ParseLevel(str string) (Level, error) {
	switch strings.ToLower(str) {
	case "debug":
		return DebugLevel, nil
	case "info":
		return InfoLevel, nil
	case "warn":
		return WarnLevel, nil
	case "error":
		return ErrorLevel, nil
	case "fatal":
		return FatalLevel, nil
	default:
		return invalid, ErrInvalidLevel
	}
}

// MarshalText implments encoding.TextMarshaller.
func (l Level) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

// UnmarshalText implments encoding.TextMarshaller.
func (l *Level) UnmarshalText(text []byte) error {
	level, err := ParseLevel(string(text))
	if err != nil {
		return err
	}

	*l = level

	return nil
}
