package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

const inputPath string = "/Users/artikow/Downloads/B-large-practice.in"
const outputPath string = "/Users/artikow/Downloads/output-revenge-of-pancakes.out"
const Eol string = "\n"

func process(b *bufio.Scanner) string {
	var s = scanText(b)
	var maxFlip int = 0

	for !AllPositive(s) {
		if AllNegative(s) {
			maxFlip++
			break
		}

		firstRune, _ := utf8.DecodeRuneInString(s)
		oFirstRune := Opponent(firstRune)
		cutIdx := strings.Index(s, oFirstRune)
		firstS := s[0:cutIdx]
		lastS := s[cutIdx:utf8.RuneCountInString(s)]
		s = Flip(firstS) + lastS
		maxFlip++

	}

	return strconv.Itoa(maxFlip)
}

func Flip(s string) string {
	var result string
	for _, v := range s {
		result = Opponent(v) + result
	}
	return result
}

func Opponent(r rune) string {
	if r == '+' {
		return "-"
	} else {
		return "+"
	}
}

func AllPositive(s string) bool {
	return allSameSign(s, '-')
}

func AllNegative(s string) bool {
	return allSameSign(s, '+')
}

func allSameSign(s string, exceptSign rune) bool {
	for _, e := range s {
		if e == exceptSign {
			return false
		}
	}
	return true
}

func main() {
	fInput, err := os.Open(inputPath)
	panicIfError(err)
	defer fInput.Close()
	fOutput, err := os.Create(outputPath)
	panicIfError(err)
	defer fOutput.Close()

	bufInput := bufio.NewScanner(fInput)
	bufOuput := bufio.NewWriter(fOutput)

	T := scanInt(bufInput)

	for i := 1; i <= T; i++ {
		result := process(bufInput)
		printResult := fmt.Sprintf("Case #%d: %s\n", i, result)
		bufOuput.WriteString(printResult)
		bufOuput.Flush()
	}
	fmt.Println("DONE")
}

func scanInt(b *bufio.Scanner) int {
	b.Scan()
	i, err := strconv.Atoi(b.Text())
	panicIfError(err)
	return i
}

func scanText(b *bufio.Scanner) string {
	b.Scan()
	return b.Text()
}

func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
}
