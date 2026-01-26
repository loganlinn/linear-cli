package format

import (
	"fmt"
	"strings"
)

// Verbosity specifies the level of detail in formatted output.
// This replaces the older Format type with clearer semantics.
type Verbosity int

const (
	// VerbosityMinimal returns only essential fields (~50 tokens per issue)
	VerbosityMinimal Verbosity = iota
	// VerbosityCompact returns commonly needed fields (~150 tokens per issue)
	VerbosityCompact
	// VerbosityFull returns all fields (~500 tokens per issue)
	VerbosityFull
)

// ParseVerbosity parses a string into a Verbosity with validation.
// Returns VerbosityCompact for empty strings (default behavior).
func ParseVerbosity(s string) (Verbosity, error) {
	if s == "" {
		return VerbosityCompact, nil // Default to compact for balanced output
	}

	switch strings.ToLower(s) {
	case "minimal", "min":
		return VerbosityMinimal, nil
	case "compact", "default":
		return VerbosityCompact, nil
	case "full", "detailed":
		return VerbosityFull, nil
	default:
		return VerbosityCompact, fmt.Errorf("invalid verbosity '%s': must be 'minimal', 'compact', or 'full'", s)
	}
}

// String returns the string representation of the verbosity level.
func (v Verbosity) String() string {
	switch v {
	case VerbosityMinimal:
		return "minimal"
	case VerbosityCompact:
		return "compact"
	case VerbosityFull:
		return "full"
	default:
		return "compact"
	}
}

// FormatToVerbosity converts the legacy Format type to Verbosity.
// This helper maintains backward compatibility during migration.
func FormatToVerbosity(format Format) Verbosity {
	switch format {
	case Minimal:
		return VerbosityMinimal
	case Full:
		return VerbosityFull
	case Compact:
		return VerbosityCompact
	default:
		return VerbosityCompact
	}
}

// VerbosityToFormat converts Verbosity back to legacy Format.
// This helper maintains backward compatibility during migration.
func VerbosityToFormat(verbosity Verbosity) Format {
	switch verbosity {
	case VerbosityMinimal:
		return Minimal
	case VerbosityFull:
		return Full
	case VerbosityCompact:
		return Compact
	default:
		return Compact
	}
}
