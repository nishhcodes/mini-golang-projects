package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var choices = []string{"rock", "paper", "scissors"}

func getComputerChoice() string {
	rand.Seed(time.Now().UnixNano())
	return choices[rand.Intn(len(choices))]
}

func getWinner(userChoice, computerChoice string) string {
	if userChoice == computerChoice {
		return "It's a tie!"
	}

	switch userChoice {
	case "rock":
		if computerChoice == "scissors" {
			return "You win! Rock crushes Scissors."
		}
		return "You lose! Paper covers Rock."

	case "paper":
		if computerChoice == "rock" {
			return "You win! Paper covers Rock."
		}
		return "You lose! Scissors cut Paper."

	case "scissors":
		if computerChoice == "paper" {
			return "You win! Scissors cut Paper."
		}
		return "You lose! Rock crushes Scissors."

	default:
		return "Invalid choice! Please choose Rock, Paper, or Scissors."
	}
}

func main() {
	var userChoice string

	fmt.Print("Enter Rock, Paper, or Scissors: ")
	fmt.Scanln(&userChoice)

	userChoice = strings.ToLower(userChoice)

	computerChoice := getComputerChoice()

	fmt.Println("Computer chose:", computerChoice)
	fmt.Println(getWinner(userChoice, computerChoice))
}
