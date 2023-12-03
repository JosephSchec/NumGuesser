package main

import (
	// "NumGuesser/functions"
	"fmt"
	"math/rand"
)

func isGreater(a, b int) bool {
	return a > b
}
func isLess(a, b int) bool {
	return a < b
}

func main() {
	fmt.Println("We have a number try and guess it")
	ourNumber := rand.Intn(9) + 1
	// fmt.Println("testNumber", ourNumber)
	// fmt.Println(functions.TestModule())
	var guess int
	for &guess != &ourNumber {
		fmt.Scanln(
			&guess,
		)
		if isGreater(guess, ourNumber) {
			fmt.Println("you guessed too high")
		} else if isLess(guess, ourNumber) {
			fmt.Println("you guessed too low")
		} else {
			break
		}
	}

	fmt.Println("YOU GOT IT")
}
