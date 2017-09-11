package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	bmath "github.com/vivalapanda/binary-math/bitmath"
)

func main() {
	stopCode := false

	fmt.Printf("Algorithms Homework 1\n")
	fmt.Printf("Usage:\n  A^B - C^D: '1'\n  A^B / C^D: '2'\n  sum(1/n, n = 1 to A): '3'\n  Kill: '4'\n")
	for stopCode == false {
		fmt.Printf("Enter your desired operation: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inText := scanner.Text()
		if len(inText) == 0 {
			continue
		}
		opChar := inText[0]

		switch opChar {
		case '1':
			fmt.Println("Please enter integers A,B,C,D in a nonspaced comma seperated list:")
			scanner.Scan()
			argText := scanner.Text()
			argArr := parseArgs(argText)
			fmt.Printf("Difference: %v\n", problem1(argArr[0], argArr[1], argArr[2], argArr[3]))
		case '2':
			fmt.Println("Please enter integers A,B,C,D in a nonspaced comma seperated list:")
			scanner.Scan()
			argText := scanner.Text()
			argArr := parseArgs(argText)
			q, r := problem2(argArr[0], argArr[1], argArr[2], argArr[3])
			fmt.Printf("Quotient: %v, Remainder: %v\n", q, r)
		case '3':
			fmt.Println("Please enter integer A")
			scanner.Scan()
			argText := scanner.Text()
			argArr := parseArgs(argText)
			num, denum := problem3(argArr[0])
			fmt.Printf("Numerator: %v, Denominator: %v\n", num, denum)
		case '4':
			fmt.Println("Shutting down")
			stopCode = true
		default:
			fmt.Println("I'm sorry, I didn't understand that...")
			fmt.Printf("Usage:\n  Kill: 'k'\n  Force Poll: 'p'\n  Debug Poll: 'd'\n")
		}
	}
}

func parseArgs(argText string) []int {
	argArr := strings.Split(argText, ",")
	intArr := make([]int, len(argArr))
	for i, elem := range argArr {
		intArr[i], _ = strconv.Atoi(elem)
		if intArr[i] > 1000 {
			intArr[i] = 1000
			fmt.Println("You entered an integer > 1000. Automatically setting to 1000")
		}
	}

	return intArr
}

func problem1(a int, b int, c int, d int) (diff int) {
	binDiff, isNegative := bmath.ProblemOne(
		bmath.Dec2bin(a),
		bmath.Dec2bin(b),
		bmath.Dec2bin(c),
		bmath.Dec2bin(d))
	diff = bmath.Bin2dec(binDiff)
	if isNegative {
		diff = -1 * diff
	}

	return diff
}

func problem2(a int, b int, c int, d int) (q int, r int) {
	binQ, binR := bmath.ProblemTwo(
		bmath.Dec2bin(a),
		bmath.Dec2bin(b),
		bmath.Dec2bin(c),
		bmath.Dec2bin(d))
	return bmath.Bin2dec(binQ), bmath.Bin2dec(binR)
}

func problem3(a int) (numerator int, denum int) {
	binNum, binDenum := bmath.ProblemThree(bmath.Dec2bin(a))
	numerator, denum = bmath.Bin2dec(binNum), bmath.Bin2dec(binDenum)

	return numerator, denum
}
