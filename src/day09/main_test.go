package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetLargestArea(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: "Ultimately, the largest rectangle you can make in this example has area 50. One way to do this is between 2,5 and 11,1:",
			inputList: []string{
				"7,1",
				"11,1",
				"11,7",
				"9,7",
				"9,5",
				"2,5",
				"2,3",
				"7,3",
			},
			expectedVal: 50,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getLargestArea(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetLargestArea2(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: "The largest rectangle you can make in this example using only red and green tiles has area 24. One way to do this is between 9,5 and 2,3:",
			inputList: []string{
				"7,1",
				"11,1",
				"11,7",
				"9,7",
				"9,5",
				"2,5",
				"2,3",
				"7,3",
			},
			expectedVal: 24,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getLargestArea2(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
