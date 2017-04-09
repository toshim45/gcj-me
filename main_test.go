package main

import "testing"

//func TestFlip(t *testing.T) {
//	raw := "++--+"
//	result := Flip(raw)
//	if result != "--++-" {
//		t.Errorf("flip failed %v -> %v\r\n", raw, result)
//	}
//}

func TestIsSorted(t *testing.T) {
	var s string
	s = "123456"
	if !IsSorted(s) {
		t.Errorf("fail %s is sorted", s)
	}
	s = "143256"
	if IsSorted(s) {
		t.Errorf("fail %s is unsorted", s)
	}
	s = "124356"
	if IsSorted(s) {
		t.Errorf("fail %s is unsorted", s)
	}
	s = "1243567"
	if IsSorted(s) {
		t.Errorf("fail %s is unsorted", s)
	}
	s = "11112223334445555666777788889999"
	if !IsSorted(s) {
		t.Errorf("fail %s is sorted", s)
	}
	s = "111111111111111110"
	if IsSorted(s) {
		t.Errorf("fail %s is unsorted", s)
	}
}

func TestFindUnsortedIdx(t *testing.T) {
	var s string
	s = "123454321"
	unsortedIdx := FindUnsortedIdx(s)
	if unsortedIdx != 5 {
		t.Error("fail expected 5 but actual", unsortedIdx)
	}
}
