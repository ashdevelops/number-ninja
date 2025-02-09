package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

type level struct {
	number    int
	name      string
	tries     int
	randomMin int
	randomMax int
}

var levels = []level{
	{
		number:    1,
		name:      "Easy",
		tries:     8,
		randomMin: 0,
		randomMax: 100,
	},
	{
		number:    2,
		name:      "Medium",
		tries:     6,
		randomMin: 1,
		randomMax: 200,
	},
	{
		number:    3,
		name:      "Hard",
		tries:     5,
		randomMin: 1,
		randomMax: 200,
	},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pick a level ")
	fmt.Println(" ")

	for i := 0; i < len(levels); i++ {
		cLevel := levels[i]

		fmt.Println(fmt.Sprintf("%d - %s - Guess a number %d to %d with %d tries",
			cLevel.number,
			cLevel.name,
			cLevel.randomMin,
			cLevel.randomMax,
			cLevel.tries))
	}

	fmt.Println(" ")

	fmt.Print("Pick a level: ")
	levelInput, _ := reader.ReadString('\n')

	levelIndex, levelInputError := strconv.Atoi(strings.Replace(levelInput, "\n", "", -1))

	if levelInputError != nil || levelIndex > len(levels) {
		panic("Invalid input")
	}

	var levelPicked = levels[levelIndex-1]

	var randomNumber = randRange(
		levelPicked.randomMin,
		levelPicked.randomMax)

	var tries = 0
	var maxTries = levelPicked.tries
	var success = false

	for i := 0; i < maxTries; i++ {
		tries++

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Guess a number between 1 and 100: ")
		text, _ := reader.ReadString('\n')

		i, err := strconv.Atoi(strings.Replace(text, "\n", "", -1))

		if err != nil {
			fmt.Println(err)
			continue
		}

		if i == randomNumber {
			success = true
			break
		} else if i < randomNumber {
			fmt.Println("Higher\n")

		} else if i > randomNumber {
			fmt.Println("Lower\n")
		}
	}

	if success {
		fmt.Println("You won!")
	} else {
		fmt.Println("You ran out of tries")
	}
}
