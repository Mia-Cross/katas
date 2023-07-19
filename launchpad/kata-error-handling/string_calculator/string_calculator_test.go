package string_calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		name     string
		given    string
		expected string
	}{
		{
			"Returns zero when given an empty string",
			"",
			"0",
		},
		{
			"Returns number when given a single number",
			"8",
			"8",
		},
		{
			"Returns the sum when given 2 positive numbers",
			"36,6",
			"42",
		},
		{
			"Returns the sum when given 5 positive numbers",
			"36,6,4,10,5",
			"61",
		},
		{
			"Returns an error when given a negative number",
			"36,6,4,-10,5",
			"Negative not allowed : -10",
		},
		{
			"Returns the sum when given 2 positive floats",
			"36.1,6.6,4,10,5.0",
			"61.7",
		},
		{
			"Displays no .0 if number is round",
			"36.1,6.6,4,10,5.0,0.3",
			"62",
		},
		{
			"Returns the sum when given '\n' as a separator",
			"36,6\n4,10\n5",
			"61",
		},
		{
			"Returns an error when given 2 separators in a row (1)",
			"36,6\n4,,10\n5",
			"Number expected but ',' found at position 7.",
		},
		{
			"Returns an error when given 2 separators in a row (2)",
			"36,6,\n4,10\n5",
			"Number expected but '\\n' found at position 5.",
		},
		{
			"Returns both errors when missing and negative values",
			"36,-6,,4,10\n\n-5",
			"Negative not allowed : -6, -5\nNumber expected but ',' found at position 6.\nNumber expected but '\\n' found at position 12.",
		},

	}
	for _, test := range tests {
		assert.Equal(t, test.expected, Add(test.given))
	}
}

/* A RAJOUTER DANS LES TESTS
	espaces, EOF, missing at begin
*/


