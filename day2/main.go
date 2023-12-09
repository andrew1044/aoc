package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Game 1: 7 red, 14 blue; 2 blue, 3 red, 3 green; 4 green, 12 blue, 15 red; 3 green, 12 blue, 3 red; 11 red, 2 green

/*
Logic

For each number and colour in each turn, add the number for the respective colur to
the currentStateMap. At the end of each turn, check each item in the currentStateMap is
less than the targetStateMap. If this conition is true, the turn is valid, get the gameNumber,
add it to the sumOFGames.
*/
func game1(stringData []string) int {
	sumOfGames := 0
	targetStateMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	// for each game
	for a, gameLineData := range stringData {
		isGameValid := 1
		gameId := a + 1

		// thanks to chatGPT - fancy way to get the data - split on ":", get the second part (discard "Game X:"),
		// then split on ";" to get the data for each turn
		turnsPerGameData := strings.Split(strings.Split(gameLineData, ":")[1], ";")

		// for each turn in game
		for _, turnData := range turnsPerGameData {
			currentStateMap := map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			}

			individualColour := strings.Split(turnData, ",")

			// for each colour in each turn
			for _, colour := range individualColour {

				// add colours per turn to the currentStateMap
				// split, then add to map -- TrimSpace removes leading and trailing spaces
				// tokenColour[0] is the integer, tokenColour[1] is the colour
				tokenPerColour := strings.Split(strings.TrimSpace(colour), " ")
				tokenCount, err := strconv.Atoi(tokenPerColour[0])
				if err != nil {
					fmt.Println("Error converting string to int" + err.Error())
					os.Exit(1)
				}
				colour := tokenPerColour[1]
				currentStateMap[colour] = tokenCount
			}

			// check current state vs target state
			for colourKey, kubeCount := range currentStateMap {
				if targetStateMap[colourKey] < kubeCount {
					isGameValid = 0
					break
				}
			}
			// ffs - need to get out of the loop
			if isGameValid == 0 {
				break
			}
		}

		if isGameValid == 1 {
			sumOfGames += gameId
		}
	}

	return sumOfGames
}

func game2(stringData []string) int {
	kubeSumPower := 0

	// for each game
	for _, gameLineData := range stringData {
		currentStateMap := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		turnsPerGameData := strings.Split(strings.Split(gameLineData, ":")[1], ";")

		for a, turnData := range turnsPerGameData {
			fmt.Print("turn: " + strconv.Itoa(a) + "\n")

			individualColour := strings.Split(turnData, ",")
			for _, colour := range individualColour {

				tokenPerColour := strings.Split(strings.TrimSpace(colour), " ")
				tokenCount, err := strconv.Atoi(tokenPerColour[0])
				if err != nil {
					fmt.Println("Error converting string to int" + err.Error())
					os.Exit(1)
				}
				tokenColour := tokenPerColour[1]

				// fmt.Print("map " + strconv.Itoa(currentStateMap[tokenColour]) + " count: " + strconv.Itoa(tokenCount) + "\n")
				// keep track of the maximum number of kubes for each colour across all turns
				// to do so, we check if the current turn value is greater than the currentStateMap
				// the result is we will have the largest red, green, blue value across all turns
				// minimum number of kubes of each colour (game objective) = max number of kubes of each colour across all turns
				if currentStateMap[tokenColour] < tokenCount {
					currentStateMap[tokenColour] = tokenCount
				}
			}

		}

		kubePower := 1
		for _, value := range currentStateMap {
			kubePower *= value
		}
		kubeSumPower += kubePower

	}

	return kubeSumPower
}

func main() {
	inputData, err := os.ReadFile("data/input_data")

	if err != nil {
		fmt.Println("Error reading file" + err.Error())
		os.Exit(1)
	}

	stringData := strings.Split(strings.TrimSpace(string(inputData)), "\n")

	// fmt.Println(stringData)

	fmt.Println("Day 2, game 1: ", game1(stringData))
	fmt.Println("Day 2, game 2: ", game2(stringData))
}
