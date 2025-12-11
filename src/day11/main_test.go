package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetNumPaths(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: "In total, there are 5 different paths leading from you to out.",
			inputList: []string{
				"aaa: you hhh",
				"you: bbb ccc",
				"bbb: ddd eee",
				"ccc: ddd eee fff",
				"ddd: ggg",
				"eee: out",
				"fff: out",
				"ggg: out",
				"hhh: ccc fff iii",
				"iii: out",
			},
			expectedVal: 5,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumPaths(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetNumPaths2(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
	}{
		{
			desc: "However, only 2 paths from svr to out visit both dac and fft.",
			inputList: []string{
				"svr: aaa bbb",
				"aaa: fft",
				"fft: ccc",
				"bbb: tty",
				"tty: ccc",
				"ccc: ddd eee",
				"ddd: hub",
				"hub: fff",
				"eee: dac",
				"dac: fff",
				"fff: ggg hhh",
				"ggg: out",
				"hhh: out",
			},
			expectedVal: 2,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumPaths2(tc.inputList)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
