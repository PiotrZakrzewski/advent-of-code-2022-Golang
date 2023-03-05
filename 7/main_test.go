package main

import "testing"

func TestOverlaping(t *testing.T) {
	a := task_range{1, 3}
	b := task_range{2, 4}
	if overlapping(a, b) {
		t.Error("Expected false")
	}
	a = task_range{1, 3}
	b = task_range{1, 3}
	if !overlapping(a, b) {
		t.Error("Expected true")
	}
}
