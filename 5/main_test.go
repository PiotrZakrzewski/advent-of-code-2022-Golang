package main

import (
	"testing"
)

// basic test for translate
func TestTranslate(t *testing.T) {
	if translate("a") != 1 {
		t.Error("Expected 1")
	}
	if translate("b") != 2 {
		t.Error("Expected 2")
	}
	if translate("z") != 26 {
		t.Error("Expected 26")
	}
	if translate("A") != 27 {
		t.Error("Expected 27")
	}
	if translate("B") != 28 {
		t.Error("Expected 28")
	}
	if translate("Z") != 52 {
		t.Error("Expected 52")
	}
}

// basic test for parse
func TestParse(t *testing.T) {
	r := parse("abAB")
	if r.compartment_1[0] != 1 {
		t.Error("Expected 1")
	}
	if r.compartment_1[1] != 2 {
		t.Error("Expected 2")
	}
	if r.compartment_2[0] != 27 {
		t.Error("Expected 27")
	}
	if r.compartment_2[1] != 28 {
		t.Error("Expected 28")
	}
}

// basic test for invalidly_packed
func TestInvalidlyPacked(t *testing.T) {
	r := parse("vJrwpWtwJgWrhcsFMMfFFhFp")
	invalid := invalidly_packed(r)
	if len(invalid) != 1 {
		t.Error("Expected 1")
	}
	if invalid[0] != 16 {
		t.Error("Expected 16")
	}
}
