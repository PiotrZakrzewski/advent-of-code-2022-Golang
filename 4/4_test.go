package main

// test result
import "testing"

func TestResult(t *testing.T) {
	got := result("Y", "A")
	if got != 4 {
		t.Error("Expected 4, got ", got)
	}
}

func TestCounter(t *testing.T) {
	got := counter("rock", "Y")
	if got != "rock" {
		t.Error("Expected rock, got ", got)
	}
}

func TestTranslate(t *testing.T) {
	got := symbol_map["A"]
	if got != "rock" {
		t.Error("Expected rock, got ", got)
	}
}
