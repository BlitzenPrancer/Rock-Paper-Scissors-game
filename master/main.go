package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
  	"math/rand"
)

var playerScore, computerScore int

func printScores(name string) {
  fmt.Printf("\n%s %d : %d %s\n\n", name, playerScore, computerScore, "computer")
}

// function to get computer choice
func getComputerChoice() rune {
  choices := []rune{'r','s','p'}
  return choices[rand.Intn(len(choices))]
}

func startGame() {
  // getting human choice
  fmt.Println("Enter r for rock")
  fmt.Println("Enter p for paper")
  fmt.Println("Enter s for scissors")
  reader := bufio.NewReader(os.Stdin)
  userChoice, _, error := reader.ReadRune()
  if error != nil {
    fmt.Println("error")
  }
  // getting computer choice
  computerChoice := getComputerChoice()
  // using maps to print rock paper scissors instead of r p s
  choices := map[string]string{"r": "rock", "p":"paper","s":"scissors"}
  fmt.Println("computer chooses",choices[string(computerChoice)])
  results(userChoice, computerChoice)
}

func main() {
	fmt.Println("......Rock Paper Scissor Game......")
	fmt.Println("Enter Your Name")
	reader := bufio.NewReader(os.Stdin) // takes I/O reader as parameter value
	// Using os package to attach standard input to that reader.
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	fmt.Printf("welcome %s \n", name)
	for {
		fmt.Println(`Press "b" to begin the game`)
		fmt.Println(`Press "e" to end the game`)
		fmt.Println(`Press "s" to see the score`)

		// Taking single character as input
		reader = bufio.NewReader(os.Stdin)
		char, _, error := reader.ReadRune()

		if error != nil {
			fmt.Println("Error occured")
			fmt.Println(error)
			break
		}

		switch char {
		case 'b':
			fmt.Println("Game starting")
      startGame()
		case 'e':
			fmt.Println("Ending Game")
      os.Exit(1)
		case 's':
			printScores(name)
		}
	}
}

func results(user rune, computer rune) {
  matchResult := string(user) + string(computer)
  switch matchResult {
    case "pr","rs","sp":
      fmt.Println("You have won the match :-)\n")
      playerScore++
    case "ps", "sr", "rp":
      fmt.Println("You have lost the match :-(\n")
      playerScore--
    case "rr","pp","ss":
      fmt.Println("Match Draw :-|\n")
  }
}
