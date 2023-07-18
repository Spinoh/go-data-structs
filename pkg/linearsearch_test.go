package day1

import "testing"

type testCase struct {
	needle   int
	expected bool
}

func TestLinearSearch(t *testing.T) {
	testArr := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	for _, tc := range getLinearTestCases() {
		result := LinearSearch(testArr, tc.needle)

		if result != tc.expected {
			t.Errorf("got: %v | expected: %v\n", result, tc.expected)
		}
	}

}

func getLinearTestCases() []testCase {
	return []testCase{
		{
			needle:   69,
			expected: true,
		},
		{
			needle:   1336,
			expected: false,
		},
		{
			needle:   69420,
			expected: true,
		},
		{
			needle:   69421,
			expected: false,
		},
		{
			needle:   1,
			expected: true,
		},
		{
			needle:   0,
			expected: false,
		},
	}
}
