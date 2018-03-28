package main

import "testing"

//Testing polymorphisms
var d1 = Polymorphism{"comp3_c0_seq1", 1349, "A", "C", "1,0", "0,0", "1,1"}

var d2 = Polymorphism{"comp3_c0_seq1", 1477, "T", "A", "1", "1,1", "0,0"}

var nd1 = Polymorphism{"comp3_c0_seq1", 1477, "T", "A", "1", "0,1", "0,0"}

//Tests
func TestDiagnosis(t *testing.T) {
	if !diagnosis(d1) {
		t.Error(`diagnosis(d1) = false`)
	}

	if !diagnosis(d2) {
		t.Error(`diagnosis(d2) = false`)
	}
}

func TestNonDiagnosis(t *testing.T) {
	if diagnosis(nd1) {
		t.Error(`diagnosis(d1) = true`)
	}
}
