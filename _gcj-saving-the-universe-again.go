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
		d, p := read(reader)
		y := "0"
		e := 0
		for d < damage(p) {
			np := swap(p)
			if strings.Compare(np, p) == 0 {
				y = "IMPOSSIBLE"
				break
			}
			p = np
			e++
			y = fmt.Sprintf("%d", e)
		}
		fmt.Printf("CASE #%d: %s\n", i+1, y)
	}
}

func swap(p string) string {
	ps := []rune(p)
	l := len(ps)
	for i := l - 1; i > 0; i-- {
		if ps[i] == 'C' {
			continue
		} else if ps[i] == 'S' && ps[i-1] == 'C' {
			ps[i] = 'C'
			ps[i-1] = 'S'
			return string(ps)
		}
	}
	return p
}

func damage(p string) int {
	c := 0
	s := 0
	for _, a := range p {
		if a == 'C' {
			c++
		} else if a == 'S' {
			s += pow2(c)
		}
	}

	return s
}

func pow2(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	d := 2
	for i := 1; i < n; i++ {
		d *= 2
	}

	return d
}

func readInt(reader *bufio.Reader) int {
	response, _, _ := reader.ReadLine()
	number, _ := strconv.Atoi(string(response))
	return number
}

func read(reader *bufio.Reader) (int, string) {
	response, _, _ := reader.ReadLine()
	words := strings.Fields(string(response))
	number, _ := strconv.Atoi(words[0])
	return number, words[1]
}

/***
d:3
p:CSCSS
dp: 2 + 4 + 4 = 10

***/
