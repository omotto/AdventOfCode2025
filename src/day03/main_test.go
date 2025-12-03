package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetSumInvalidIDs(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		digits      int
		expectedVal int
	}{
		{
			desc: "The total output joltage is the sum of the maximum joltage from each bank, so in this example, the total output joltage is 98 + 89 + 78 + 92 = 357.",
			inputList: []string{
				"987654321111111",
				"811111111111119",
				"234234234234278",
				"818181911112111",
			},
			digits:      2,
			expectedVal: 357,
		},
		{
			desc: "The total output joltage is now much larger: 987654321111 + 811111111119 + 434234234278 + 888911112111 = 3121910778619.",
			inputList: []string{
				"987654321111111",
				"811111111111119",
				"234234234234278",
				"818181911112111",
			},
			digits:      12,
			expectedVal: 3121910778619,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getSumJoltage(tc.inputList, tc.digits)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
