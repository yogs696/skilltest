package usecase

import (
	"strings"
)

// ParseUnwantedError parse and delete unwanted error message
func ParseUnwantedError(err error) string {
	if err == nil {
		return ""
	}

	var s string
	estr := err.Error()

	switch true {
	case strings.Contains(estr, "context deadline exceeded"):
		s = "context deadline exceeded, connection timeout"

	case strings.Contains(estr, "Timeout exceeded"):
		s = "dial tcp error timeout"

	case strings.Contains(estr, "SQLSTATE 42P01"):
		s = "desired table does not exist error"

	case strings.Contains(estr, "SQLSTATE 42703"):
		s = "desired column does not exist error"

	case strings.Contains(estr, "SQLSTATE 23502"):
		s = "some column value cannot be null"

	case strings.Contains(estr, "SQLSTATE 22001"):
		s = "value too long"

	case strings.Contains(estr, "SQLSTATE 23503"):
		s = "violates foreign key constraint"

	default:
		s = estr
	}

	return s
}
