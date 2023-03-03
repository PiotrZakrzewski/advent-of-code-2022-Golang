package main

// test result
import "testing"

func TestResult(t *testing.T) {
	got := result("Y", "A")
	if got != 8 {
		t.Error("Expected 8, got ", got)
	}
}

func TestTranslate(t *testing.T) {
	got := symbol_map["A"]
	if got != "rock" {
		t.Error("Expected rock, got ", got)
	}
}
