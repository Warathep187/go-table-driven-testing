package main

import (
	"testing"
)

func TestSum_EmptySlice(t *testing.T) {
	got := Sum()
	expected := 0
	if got != expected {
		t.Errorf("Sum() = %d; want %d", got, expected)
	}
}

func TestSum_SingleNumber(t *testing.T) {
	got := Sum(1)
	expected := 1
	if got != expected {
		t.Errorf("Sum(1) = %d; want %d", got, expected)
	}
}

func TestSum_MultipleNumbers(t *testing.T) {
	got := Sum(1, 2, 3, 4, 5)
	expected := 15
	if got != expected {
		t.Errorf("Sum(1, 2, 3, 4, 5) = %d; want %d", got, expected)
	}
}

func TestSum_NegativeNumbers(t *testing.T) {
	got := Sum(-1, -2, -3, -4, -5)
	expected := -15
	if got != expected {
		t.Errorf("Sum(-1, -2, -3, -4, -5) = %d; want %d", got, expected)
	}
}
