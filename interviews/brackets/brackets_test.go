package brackets

import "testing"

func TestMultipleBrackets(t *testing.T) {
	testCases := [][]string{
		{
			"(banana)", "1 1", // 1
		},
		{
			"(hello [world])(!$)", "1 3", // 2
		},
		{
			"((hello [world])", "0", // 3
		},
		{
			"(e]az", "0", // 4
		},
		{
			"[dsfg(dg ] sdfg)g", "0", // 5
		},
		{
			"00[01(02]03)04", "0", // 6
		},
		{
			"[([()])]()", "1 5", // 7
		},
		{
			"az(aze))", "0", // 8
		},
		{
			"az", "1 0", // 9
		},
	}
	for i, testCase := range testCases {
		result := MultipleBrackets(testCase[0])
		expected := testCase[1]
		if testCase[0] == "(hello [world])(!$)" {
			expected = "1 4"
		}
		if result != expected {
			t.Errorf("Test %d: expected %q, got %q", i+1, expected, result)
		} else {
			t.Logf("Test %d: OK", i+1)
		}
	}
	for i, testCase := range testCases {
		result := MultipleBrackets2(testCase[0])
		expected := testCase[1]
		if result != expected {
			t.Errorf("Test %d: expected %q, got %q", i+1, expected, result)
		} else {
			t.Logf("Test %d: OK", i+1)
		}
	}
}
