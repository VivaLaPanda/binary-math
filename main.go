package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
			fmt.Printf("Usage:\n  A^B - C^D: '1'\n  A^B / C^D: '2'\n  sum(1/n, n = 1 to A): '3'\n  Kill: '4'\n")
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
