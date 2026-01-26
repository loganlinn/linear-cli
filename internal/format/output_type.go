package format

import (
	"fmt"
	"strings"
)

// OutputType specifies the renderer to use for formatting output.
type OutputType string

const (
	// OutputText renders output as ASCII text (default, token-efficient)
	OutputText OutputType = "text"
	// OutputJSON renders output as JSON (machine-readable)
	OutputJSON OutputType = "json"
)

// ParseOutputType parses a string into an OutputType with validation.
// Returns OutputText for empty strings (default behavior).
func ParseOutputType(s string) (OutputType, error) {
	if s == "" {
		return OutputText, nil // Default to text for backward compatibility
	}

	switch strings.ToLower(s) {
	case "text", "ascii", "txt":
		return OutputText, nil
	case "json":
		return OutputJSON, nil
	default:
		return OutputText, fmt.Errorf("invalid output type '%s': must be 'text' or 'json'", s)
	}
}

// String returns the string representation of the output type.
func (o OutputType) String() string {
	return string(o)
}

// IsJSON returns true if this output type is JSON.
func (o OutputType) IsJSON() bool {
	return o == OutputJSON
}

// IsText returns true if this output type is text.
func (o OutputType) IsText() bool {
	return o == OutputText
}
