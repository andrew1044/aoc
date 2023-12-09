package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func firstDigit(line string) int {
	for _, c := range line {
		if unicode.IsDigit(c) {
			return int(c - '0')
		}
	}
	return 0
}

func lastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			return int(line[i] - '0')
		}
	}
	return 0
}

func convertTerms(line string) (out string) {
	conversionMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	finalIntegers := []int{}
	for charIndex, char := range line {
		if digit, err := strconv.Atoi(string(char)); err == nil {
			finalIntegers = append(finalIntegers, digit)
		} else {
			for strNumber, number := range conversionMap {
				// line[charIndex:] is a substring starting at charIndex onwards
				// strNumber represents the number as a string in the map
				// so this is "if the following characters from the index match the key in the map"
				if strings.HasPrefix(line[charIndex:], strNumber) {
					finalIntegers = append(finalIntegers, number)
					break
				}
			}
		}
	}

	// above produces a '[]int{}', convent to a string so the firstDigit and lastDigit functions can be used
	strDigits := make([]string, len(finalIntegers))
	for i, digit := range finalIntegers {
		strDigits[i] = strconv.Itoa(digit)
	}

	return strings.Join(strDigits, "")
}

func partTwo(filename string) (total int) {

	ReadFile, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer ReadFile.Close()

	fileScanner := bufio.NewScanner(ReadFile)

	sum := 0

	for fileScanner.Scan() {
		// line := fileScanner.Text()
		convertedLine := convertTerms(fileScanner.Text())
		// fmt.Print(line + " ---> " + convertedLine + "\n")

		firstDig := firstDigit(convertedLine)
		lastDig := lastDigit(convertedLine)

		result := firstDig*10 + lastDig
		// fmt.Println("result: ", result)
		sum += result
	}

	// fmt.Println("Sum: ", sum)
	return sum
}

func partOne(filename string) (total int) {
	ReadFile, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer ReadFile.Close()

	fileScanner := bufio.NewScanner(ReadFile)

	sum := 0
	index := 0
	for fileScanner.Scan() {
		// fmt.Println(strconv.Itoa(index) + " " + fileScanner.Text())
		firstDig := firstDigit(fileScanner.Text())
		lastDig := lastDigit(fileScanner.Text())

		result := firstDig*10 + lastDig
		sum += result

		// fmt.Println(strconv.Itoa(result))
		index++
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	// first half of question
	// fmt.Println("Sum: ", sum)
	return sum
}

func main() {
	p1 := partOne("data/input_data")
	fmt.Println("Part One Result: ", p1)
	p2 := partTwo("data/input_data")
	fmt.Println("Part Two Result: ", p2)
}
