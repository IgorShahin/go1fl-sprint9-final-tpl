package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	nums := generateRandomElements(10)

	if len(nums) != 10 {
		t.Error("Expected 10 elements, got ", len(nums))
	}

	for _, num := range nums {
		if num < 0 {
			t.Error("Expected non-negative numbers, got ", num)
		}
	}

	if len(generateRandomElements(0)) != 0 {
		t.Error("expected empty slice for size 0")
	}

	if len(generateRandomElements(-5)) != 0 {
		t.Error("expected empty slice for negative size")
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		data []int
		want int
	}{
		{[]int{}, 0},
		{[]int{5}, 5},
		{[]int{1, 9, 3}, 9},
		{[]int{-5, -2, -8}, -2},
	}

	for _, test := range tests {
		if got := maximum(test.data); got != test.want {
			t.Errorf("maximum(%v) = %d, want %d", test.data, got, test.want)
		}
	}
}

func TestMacChunks(t *testing.T) {
	tests := [][]int{
		{},
		{5},
		{1, 9, 3},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{-5, -2, -8},
	}

	for _, test := range tests {
		if got, want := maxChunks(test), maximum(test); got != want {
			t.Errorf("maxChunks(%v) = %d, want %d", test, got, want)
		}
	}
}
