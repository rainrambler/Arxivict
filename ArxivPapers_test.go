package main

import (
	"testing"
)

func TestGetCategory1(t *testing.T) {
	s := "[cs.GL]"
	res := getCategory(s)
	expected := "[cs]"

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestGetCategory2(t *testing.T) {
	s := "[nucl-th]"
	res := getCategory(s)
	expected := s

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}
