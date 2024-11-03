package main

import (
	"testing"
)

func TestArrToJsonStr1(t *testing.T) {
	arr := []int{200, 300, 40}
	res := ArrToJsonStr(arr)
	expected := "[200, 300, 40]"

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}
