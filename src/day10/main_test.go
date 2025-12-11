package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetFewestButtonClicks(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: "So, the fewest button presses required to correctly configure the indicator lights on all of the machines is 2 + 3 + 2 = 7.",
			inputList: []string{
				"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
				"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
				"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
			},
			expectedVal: 7,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getFewestButtonClicks(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetFewestButtonClicks2(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: "So, the fewest button presses required to correctly configure the joltage level counters on all of the machines is 10 + 12 + 11 = 33.",
			inputList: []string{
				"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
				"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
				"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
			},
			expectedVal: 33,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getFewestButtonClicks2(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
