package main

import "testing"

func TestPow2(t *testing.T) {
	var r int
	if r = pow2(0); r != 1 {
		t.Error("should 1 but: ", r)
	}
	if r = pow2(1); r != 2 {
		t.Error("should 2 but: ", r)
	}
	if r = pow2(2); r != 4 {
		t.Error("should 4 but: ", r)
	}
	if r = pow2(3); r != 8 {
		t.Error("should 8 but: ", r)
	}
	if r = pow2(30); r != 1073741824 {
		t.Error("should 1073741824 but: ", r)
	}
}

func TestDamage(t *testing.T) {
	var r int
	if r = damage("SS"); r != 2 {
		t.Error("should 2 but: ", r)
	}
	if r = damage("CC"); r != 0 {
		t.Error("should 0 but: ", r)
	}
	if r = damage("CCCCCCCCCCCCCCCCCCCCCCCCCCCCCC"); r != 0 {
		t.Error("should 0 but: ", r)
	}
	if r = damage("SCCSSC"); r != 9 {
		t.Error("should 9 but: ", r)
	}
	if r = damage("SCSCSC"); r != 7 {
		t.Error("should 7 but: ", r)
	}
	if r = damage("SCSSCC"); r != 5 {
		t.Error("should 5 but: ", r)
	}
}

func TestSwap(t *testing.T) {
	var r string
	if r = swap("SC"); r != "SC" {
		t.Error("should SC but: ", r)
	}
	if r = swap("CS"); r != "SC" {
		t.Error("should SC but: ", r)
	}
	if r = swap("CC"); r != "CC" {
		t.Error("should CC but: ", r)
	}
	if r = swap("SS"); r != "SS" {
		t.Error("should SS but: ", r)
	}
	if r = swap("CSS"); r != "SCS" {
		t.Error("should SCS but: ", r)
	}
	if r = swap("SCCSSC"); r != "SCSCSC" {
		t.Error("should SCSCSC but: ", r)
	}
	if r = swap("SCSCSC"); r != "SCSSCC" {
		t.Error("should SCSSCC but: ", r)
	}
}
