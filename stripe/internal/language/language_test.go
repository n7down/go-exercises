package language

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAcceptLanguage(t *testing.T) {
	tests := []struct {
		header   string
		accepted []string
		expected []string
	}{
		{
			header:   "",
			accepted: []string{"en-US"},
			expected: []string{},
		}, {
			header:   "en-US",
			accepted: []string{},
			expected: []string{},
		}, {
			header:   "en-US, en-US",
			accepted: []string{},
			expected: []string{},
		}, {
			header:   "en-US, en-US",
			accepted: []string{"en-US"},
			expected: []string{"en-US"},
		}, {
			header:   "en-US, en-US, fr-CA",
			accepted: []string{"en-US", "fr-CA"},
			expected: []string{"en-US", "fr-CA"},
		}, {
			header:   "en-US, fr-CA, fr-FR",
			accepted: []string{"fr-FR", "en-US"},
			expected: []string{"en-US", "fr-FR"},
		}, {
			header:   "fr-CA, fr-FR",
			accepted: []string{"en-US", "fr-FR"},
			expected: []string{"fr-FR"},
		}, {
			header:   "en",
			accepted: []string{"en-US", "fr-CA", "fr-FR"},
			expected: []string{"en-US"},
		},
	}

	for _, test := range tests {
		actual := ParseAcceptLanguage(test.header, test.accepted)
		assert.Equal(t, test.expected, actual)
	}
}
