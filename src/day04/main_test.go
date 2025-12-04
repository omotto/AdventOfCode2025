package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetSumRolls(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: "In this example, there are 13 rolls of paper that can be accessed by a forklift.",
			inputList: []string{
				"..@@.@@@@.",
				"@@@.@.@.@@",
				"@@@@@.@.@@",
				"@.@@@@..@.",
				"@@.@@@@.@@",
				".@@@@@@@.@",
				".@.@.@.@@@",
				"@.@@@.@@@@",
				".@@@@@@@@.",
				"@.@.@@@.@.",
			},
			expectedVal: 13,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getSumRolls(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetSumRolls2(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: "Stop once no more rolls of paper are accessible by a forklift. In this example, a total of 43 rolls of paper can be removed.",
			inputList: []string{
				"..@@.@@@@.",
				"@@@.@.@.@@",
				"@@@@@.@.@@",
				"@.@@@@..@.",
				"@@.@@@@.@@",
				".@@@@@@@.@",
				".@.@.@.@@@",
				"@.@@@.@@@@",
				".@@@@@@@@.",
				"@.@.@@@.@.",
			},
			expectedVal: 43,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getSumRolls2(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
