package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetSumOperations(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: "In this worksheet, the grand total is 33210 + 490 + 4243455 + 401 = 4277556.",
			inputList: []string{
				"123 328  51 64",
				"45 64  387 23",
				"6 98  215 314",
				"*   +   *   +",
			},
			expectedVal: 4277556,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getSumOperations(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetSumOperations2(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: "Now, the grand total is 1058 + 3253600 + 625 + 8544 = 3263827.",
			inputList: []string{
				"123 328  51 64 ",
				" 45 64  387 23 ",
				"  6 98  215 314",
				"*   +   *   +  ",
			},
			expectedVal: 3263827,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getSumOperations2(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
