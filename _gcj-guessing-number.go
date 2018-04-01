package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	CORRECT     = "CORRECT"
	WRONGANSWER = "WRONG_ANSWER"
	TOOBIG      = "TOO_BIG"
	TOOSMALL    = "TOO_SMALL"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	t := readInt(reader)
	for i := 0; i < t; i++ {
		a, b := readTwoInt(reader)
		n := readInt(reader)
		d := n
		g := (b - a) / 2

		fmt.Println(g)
		r := read(reader)
		if r == CORRECT {
			continue
		}

		pa := r

		for r != CORRECT {
			if r == TOOBIG {
				if pa == TOOSMALL {
					d /= 2
					if d == 0 {
						d = 1
					}
				}
				g -= d
				if g <= a {
					g = a + 1
				}
			} else if r == TOOSMALL {
				if pa == TOOBIG {
					d /= 2
					if d == 0 {
						d = 1
					}
				}
				g += d
				if g > b {
					g = b - 1
				}
			} else if r == WRONGANSWER {
				os.Exit(1)
			}
			fmt.Println(g)
			r = read(reader)
		}
	}
}

func read(reader *bufio.Reader) string {
	response, _, _ := reader.ReadLine()
	return string(response)
}

func readInt(reader *bufio.Reader) int {
	response, _, _ := reader.ReadLine()
	number, _ := strconv.Atoi(string(response))
	return number
}

func readTwoInt(reader *bufio.Reader) (int, int) {
	response, _, _ := reader.ReadLine()
	numbers := strings.Fields(string(response))
	number0, _ := strconv.Atoi(numbers[0])
	number1, _ := strconv.Atoi(numbers[1])
	return number0, number1
}

/***

1 2 3 4 5 6 7 8 9 10

10 -> too_big
10 - 10 = 1 -> too_small
1 + 5 -> 6 -> too_big
6 - (2) = 4 -> too_small
4 + (1) = 5 -> correct


***/
