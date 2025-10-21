package helper

import (
	"time"
)

// DateStrToUnixNano convert date string formated to epoch unix nano seconds.
func DateStrToUnixNano(value string, layout ...string) int64 {
	format := "2006-01-02"
	if len(layout) > 0 {
		format = layout[0]
	}

	t, err := time.Parse(format, value)
	if err != nil {
		return 0
	}

	return t.UnixNano()
}

// DateStrToUnixNanoStrict same like DateStrToUnixNano, but will return error if got any error.
func DateStrToUnixNanoStrict(value string, layout ...string) (int64, error) {
	format := "2006-01-02"
	if len(layout) > 0 {
		format = layout[0]
	}

	t, err := time.Parse(format, value)
	if err != nil {
		return 0, err
	}

	return t.UnixNano(), nil
}
