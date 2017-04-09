package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	inputPath  string = "/Users/artikow/Downloads/B-large.in"
	outputPath string = "/Users/artikow/Downloads/tidy-number.out"
	Eol        string = "\n"
)

func main() {
	start := time.Now()
	var input string
	if len(os.Args) > 1 {
		input = os.Args[1]
	} else {
		input = inputPath
	}
	fInput, err := os.Open(input)
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
		fmt.Printf(printResult)
		bufOuput.WriteString(printResult)
		bufOuput.Flush()
	}
	fmt.Println(">> DONE ", time.Since(start), " <<")
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

func process(b *bufio.Scanner) string {
	var num = scanInt(b)
	var isTidy bool
	for !isTidy && num > 0 {
		isTidy, num = CalculateTidy(num)
		if !isTidy {
			num--
		}
	}
	return strconv.Itoa(num)
}

func CalculateTidy(num int) (bool, int) {
	if num%10 == 0 {
		return false, num
	}

	if num < 10 {
		return true, num
	}

	var numText = strconv.Itoa(num)

	isSorted := IsSorted(numText)

	unsortedIdx := FindUnsortedIdx(numText)

	if unsortedIdx != len(numText) && !isSorted {
		zeroed, _ := strconv.Atoi(numText[unsortedIdx:])
		num = num - zeroed
	}

	return isSorted, num
}

func FindUnsortedIdx(s string) int {
	sLen := len(s)
	for i := 0; i < sLen-1; i++ {
		if s[i] > s[i+1] {
			return i + 1
		}
	}
	return sLen
}

func IsSorted(s string) bool {
	sLen := len(s)

	if sLen == 2 {
		return s[0] <= s[1]
	}

	if sLen == 3 {
		return s[0] <= s[1] && s[1] <= s[2]
	}

	return IsSorted(s[:sLen/2]) && IsSorted(s[sLen/2-1:sLen/2+1]) && IsSorted(s[sLen/2:])
}

/**
Problem

Tatiana likes to keep things tidy. Her toys are sorted from smallest to largest,
her pencils are sorted from shortest to longest and her computers from oldest to newest.
One day, when practicing her counting skills, she noticed that some integers,
when written in base 10 with no leading zeroes, have their digits sorted in non-decreasing order.
Some examples of this are 8, 123, 555, and 224488. She decided to call these numbers tidy.
Numbers that do not have this property, like 20, 321, 495 and 999990, are not tidy.

She just finished counting all positive integers in ascending order from 1 to N.
What was the last tidy number she counted?

Input

The first line of the input gives the number of test cases, T. T lines follow.
Each line describes a test case with a single integer N, the last number counted by Tatiana.
Output

For each test case, output one line containing Case #x: y,
where x is the test case number (starting from 1) and y is the last tidy number counted by Tatiana.
Limits

1 ≤ T ≤ 100.
Small dataset

1 ≤ N ≤ 1000.
Large dataset

1 ≤ N ≤ 1018.
Sample

Input

Output


4
132
1000
7
111111111111111110



Case #1: 129
Case #2: 999
Case #3: 7
Case #4: 99999999999999999

Note that the last sample case would not appear in the Small dataset.

**/
