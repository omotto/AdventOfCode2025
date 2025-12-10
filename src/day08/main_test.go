package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetProductCircuits(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: "Multiplying together the sizes of the three largest circuits (5, 4, and one of the circuits of size 2) produces 40.",
			inputList: []string{
				"162,817,812",
				"57,618,57",
				"906,360,560",
				"592,479,940",
				"352,342,300",
				"466,668,158",
				"542,29,236",
				"431,825,988",
				"739,650,466",
				"52,470,668",
				"216,146,977",
				"819,987,18",
				"117,168,530",
				"805,96,715",
				"346,949,466",
				"970,615,88",
				"941,993,340",
				"862,61,35",
				"984,92,344",
				"425,690,689",
			},
			expectedVal: 40,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getProductCircuits(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetProductCircuits2(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: "Multiplying the X coordinates of those two junction boxes (216 and 117) produces 25272.",
			inputList: []string{
				"162,817,812",
				"57,618,57",
				"906,360,560",
				"592,479,940",
				"352,342,300",
				"466,668,158",
				"542,29,236",
				"431,825,988",
				"739,650,466",
				"52,470,668",
				"216,146,977",
				"819,987,18",
				"117,168,530",
				"805,96,715",
				"346,949,466",
				"970,615,88",
				"941,993,340",
				"862,61,35",
				"984,92,344",
				"425,690,689",
			},
			expectedVal: 25272,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getProductCircuits2(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
