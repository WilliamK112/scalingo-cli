package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateHoursBeforeDelete(t *testing.T) {
	validate := validateHoursBeforeDelete(t.Context())

	for _, value := range []string{"", "0", "1", "2147483647"} {
		t.Run("accepts "+value, func(t *testing.T) {
			require.NoError(t, validate(value))
		})
	}

	tests := map[string]struct {
		value string
		error string
	}{
		"negative":     {value: "-1", error: "must be positive"},
		"not a number": {value: "later", error: "error parsing hours"},
		"out of range": {value: "2147483648", error: "error parsing hours"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			require.ErrorContains(t, validate(test.value), test.error)
		})
	}
}

func TestParseHoursBeforeDelete(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected uint
	}{
		"empty uses default": {input: "", expected: 0},
		"zero":               {input: "0", expected: 0},
		"positive":           {input: "12", expected: 12},
		"maximum":            {input: "2147483647", expected: 2147483647},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := parseHoursBeforeDelete(t.Context(), test.input)
			require.NoError(t, err)
			require.Equal(t, test.expected, actual)
		})
	}
}
