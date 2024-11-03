package main

import (
	"testing"
)

func TestGetYears1(t *testing.T) {
	ymin := 2002
	ymax := 2005
	years := GetYears(ymin, ymax)
	res := len(years)
	expected := 4

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}
