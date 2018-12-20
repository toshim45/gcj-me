package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	t := readInt(reader)
	for i := 0; i < t; i++ {
		y := "0"
		n := readInt(reader)
		v := readIntList(reader, n)
		r := troubleSort(v, n)
		rOk, rIdx := checkOrder(r, n)
		if rOk {
			y = "OK"
		} else {
			y = fmt.Sprintf("%d", rIdx)
		}
		fmt.Printf("CASE #%d: %s\n", i+1, y)
	}
}

func troubleSort(vs []int, l int) []int {
	done := false
	for !done {
		done = true
		for i := 0; i < l-2; i++ {
			if vs[i] > vs[i+2] {
				done = false
				temp := vs[i]
				vs[i] = vs[i+2]
				vs[i+2] = temp
			}
		}
	}
	return vs
}

func checkOrder(vs []int, l int) (bool, int) {
	for i := 0; i < l-1; i++ {
		if vs[i] > vs[i+1] {
			return false, i
		}
	}
	return true, 0
}

func sliceIntEqual(s1 []int, s2 []int) bool {
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func readInt(reader *bufio.Reader) int {
	response, _, _ := reader.ReadLine()
	number, _ := strconv.Atoi(string(response))
	return number
}

func readIntList(reader *bufio.Reader, n int) []int {
	response, _, _ := reader.ReadLine()
	numbers := strings.Fields(string(response))
	r := make([]int, n)
	for i, number := range numbers {
		r[i], _ = strconv.Atoi(number)
	}

	return r
}
