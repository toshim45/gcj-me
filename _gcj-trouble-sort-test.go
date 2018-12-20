package main

import "testing"

func TestTroubleSort(t *testing.T) {
	var r, i, o []int
	i = []int{5, 6, 6, 4, 3}
	o = []int{3, 4, 5, 6, 6}
	if r = troubleSort(i, len(i)); !sliceIntEqual(r, o) {
		t.Error("should be ", r, " but: ", o)
	}

	i = []int{8, 9, 7}
	o = []int{7, 9, 8}
	if r = troubleSort(i, len(i)); !sliceIntEqual(r, o) {
		t.Error("should be ", r, " but: ", o)
	}
}

func TestCheckOrder(t *testing.T) {
	var rOk bool
	var rIdx int

	if rOk, rIdx = checkOrder([]int{1, 2, 3, 4, 5}, 5); !rOk || rIdx != 0 {
		t.Error("should true, 0 but: ", rOk, ",", rIdx)
	}

	if rOk, rIdx = checkOrder([]int{8, 9, 7}, 3); rOk || rIdx != 1 {
		t.Error("should false, 1 but: ", rOk, ",", rIdx)
	}

	if rOk, rIdx = checkOrder([]int{1, 3, 2, 4}, 4); rOk || rIdx != 1 {
		t.Error("should false, 2 but: ", rOk, ",", rIdx)
	}

	if rOk, rIdx = checkOrder([]int{1, 3, 2}, 3); rOk || rIdx != 1 {
		t.Error("should false, 2 but: ", rOk, ",", rIdx)
	}
}
