package game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"../model"
)

// total number of rounds to be played
const NumberOfRounds int = 5
const check = 0

func DiceGame(players []model.Players) ([]string, int) {

	min := 1
	max := 6
	// getting different randon number every time.
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < NumberOfRounds; i++ {
		for j := 0; j < len(players); j++ {
			digit := 0
			inst := ""
			fmt.Print(players[j].Name)
			// asking every player to roll the device
			fmt.Println("  Please.... Press r to Roll the dice")
			fmt.Scanln(&inst)

			if strings.ToLower(inst) == "r" {
				// generating different random  number between 1 to 6
				digit = rand.Intn(max-min) + min
				fmt.Println("Your dice shows: ", digit)
				if digit == 1 || digit == 3 || digit == 5 {
					players[j].Score = players[j].Score + 5
				}
				if digit == 2 || digit == 4 || digit == 6 {
					players[j].Score = players[j].Score - 3
				}
				players[j].Counter = players[j].Counter + 1
				fmt.Println("\n")

			} else {
				fmt.Println(" You have pressed wrong key...Please press the right key next time ")

			}
		}
	}
	// find maximum score from all players and correspanding players name
	winnerScore := players[0].Score
	var winnerName []string
	for j := 0; j < len(players); j++ {

		if players[j].Score >= winnerScore {
			winnerScore = players[j].Score
		}
	}
	for j := 0; j < len(players); j++ {

		if players[j].Score == winnerScore {
			winnerName = append(winnerName, players[j].Name)
		}
	}

	// printing score card
	input := "Name"
	input2 := "Score"
	padded1 := fmt.Sprintf("%5v", input)
	padded2 := fmt.Sprintf("%9v", input2)
	fmt.Println("Here is the Score Card")
	fmt.Printf(padded1)
	fmt.Printf(padded2)
	fmt.Println("")
	for j := 0; j < len(players); j++ {
		fmt.Printf("%6v",players[j].Name)
		fmt.Printf("%8v", players[j].Score)
		fmt.Println("")
	}
	//  returning winner score and winner name to main file
	return winnerName, winnerScore
}
