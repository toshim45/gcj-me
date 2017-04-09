package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputPath string = "/Users/artikow/Downloads/A-large.in"
const outputPath string = "/Users/artikow/Downloads/output-counting-sheep.out"
const Eol string = "\n"
const Insomnia string = "INSOMNIA"

func process(b *bufio.Scanner) string {
	var n = scanInt(b)
	var lastN int
	if n == 0 {
		return Insomnia
	}
	tenDigit := map[rune]int{'0': 0, '1': 0, '2': 0, '3': 0, '4': 0, '5': 0, '6': 0, '7': 0, '8': 0, '9': 0}

	for i := 1; tenDigitStillContainZero(tenDigit); i++ {
		lastN = n * i
		verifyTenDigit(tenDigit, lastN)
	}
	return strconv.Itoa(lastN)
}

func verifyTenDigit(m map[rune]int, n int) {
	s := strconv.Itoa(n)
	for _, c := range s {
		m[c] = 1
	}
}

func tenDigitStillContainZero(m map[rune]int) bool {
	return m['0'] == 0 || m['1'] == 0 ||
		m['2'] == 0 || m['3'] == 0 ||
		m['4'] == 0 || m['5'] == 0 ||
		m['6'] == 0 || m['7'] == 0 ||
		m['8'] == 0 || m['9'] == 0
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
