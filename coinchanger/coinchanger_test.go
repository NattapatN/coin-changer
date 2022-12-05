package coinchanger

import (
	"testing"
)

func TestCoinChanger(t *testing.T) {
	testCases := []struct {
		input          int
		expectedOutput []int
	}{
		{10, []int{10}},
		{13, []int{10, 2, 1}},
		{25, []int{10, 10, 5}},
	}

	for _, tc := range testCases {
		output := coinchanger(tc.input)
		if !compareArray(output, tc.expectedOutput) {
			t.Errorf("incorrect output for `%#v`: expected `%v` but got `%v`", tc.input, tc.expectedOutput, output)
		}
	}
}

func compareArray(output, expectOutput []int) bool {
	if len(output) != len(expectOutput) {
		return false
	}
	for i, v := range output {
		if v != expectOutput[i] {
			return false
		}
	}
	return true
}
