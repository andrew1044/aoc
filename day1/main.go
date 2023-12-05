package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func main() {
	ReadFile, err := os.Open("input_data")

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

		fmt.Println(strconv.Itoa(result))
		index++
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	// first half of question
	fmt.Println("Sum: ", sum)

}
