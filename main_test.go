package main

import (
	"testing"
)

func TestIsHeavy(t *testing.T) {
	type testcase struct {
		p        Package
		expected bool
	}
	cases := []testcase{
		{Package{0, 0, 0, 1}, false},
		{Package{0, 0, 0, HeavyWeightThreshold - 1}, false},
		{Package{0, 0, 0, HeavyWeightThreshold}, true},
		{Package{0, 0, 0, HeavyWeightThreshold + 1}, true},
	}
	for _, tc := range cases {
		got := tc.p.IsHeavy()
		if got != tc.expected {
			t.Error("For", tc.p, "expected", tc.expected, "got", got)
		}
	}
}

func TestIsBulky(t *testing.T) {
	type testcase struct {
		p        Package
		expected bool
	}
	cases := []testcase{
		{Package{1, 1, 1, 1}, false},
		{Package{BulkySpatialThreshold - 1, 1, 1, 1}, false},
		{Package{BulkySpatialThreshold, 1, 1, 1}, true},
		{Package{1, BulkySpatialThreshold, 1, 1}, true},
		{Package{1, 1, BulkySpatialThreshold, 1}, true},
		{Package{100, 100, 99, 1}, false},
		{Package{100, 100, 100, 1}, true},
	}
	for _, tc := range cases {
		got := tc.p.IsBulky()
		if got != tc.expected {
			t.Error("For", tc.p, "expected", tc.expected, "got", got)
		}
	}
}

func TestSort(t *testing.T) {
	type testcase struct {
		p        Package
		expected Classification
	}
	cases := []testcase{
		{Package{1, 2, 3, 4}, Standard},                          // not heavy, not bulky
		{Package{1, 1, 1, HeavyWeightThreshold}, Special},        // heavy, not bulky
		{Package{100, 100, 100, 1}, Special},                     // not heavy, bulky
		{Package{100, 100, 100, HeavyWeightThreshold}, Rejected}, // heavy, bulky
	}
	for _, tc := range cases {
		got := tc.p.Sort()
		if got != tc.expected {
			t.Error("For", tc.p, "expected", tc.expected, "got", got)
		}
	}
}
