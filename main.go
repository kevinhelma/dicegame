package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./game"

	"./model"
)

func clearScreen(){
	for i:=0; i<10; i++{
		fmt.Println("\n")
	}
}

func main() {
	var players []model.Players
	var numberOfPlayers int
	var winnerName []string
	var winnerScore int
	var name string
	var confirmationCall string
	reader := bufio.NewReader(os.Stdin)

	// asking number of players want to play
	fmt.Print("Enter number of players -->  ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.TrimSpace(strings.Replace(text, "\n", "", -1))
	i, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Please enter a valid value", err)
	} else {
		numberOfPlayers = i
	}
	playersName := make([]string, numberOfPlayers)

	if numberOfPlayers >= 2 {
		// taking players name as input
		for i := 0; i < numberOfPlayers; i++ {
			message := fmt.Sprintf("Enter player %d name", i+1)
			fmt.Println(message)
			fmt.Scanln(&name)
			fmt.Println("Player name is:", name)
			playersName[i] = name
			clearScreen()
		}

		// showing all players who want to play
		for s := 0; s < len(playersName); s++ {
			fmt.Print("Player ", s+1, "-->  ")
			fmt.Println(playersName[s])
			p := model.Players{Name: playersName[s], Counter: 0, Score: 0}
			players = append(players, p)
		}

		fmt.Println("Lets Start the Game.....GOOD LUCK")
		// asking for confirmation whether he want to play or not
		fmt.Println("Please.... Press c to Continue the game")
		fmt.Scanln(&confirmationCall)
		fmt.Println("\n")

		// call the game functionalty file and will get back winner score and winner name
		if strings.ToLower(confirmationCall) == "c" {
			winnerName, winnerScore = game.DiceGame(players)

		} else {
			fmt.Println("Thank you for your time.....Have a great day...|||")
		}
		// printing winner name and their score card
		if len(winnerName) != 0 {
			if len(winnerName) == 1 {
				fmt.Print("Congratulation...Winner is ")
				fmt.Println(winnerName[0])
			} else if len(winnerName) > 1 {
				fmt.Print("Match is Tied between ")

				for j := 0; j < len(winnerName); j++ {

					// fmt.Print(winnerName[j], " ")

					if j == len(winnerName)-1 {
						fmt.Print(winnerName[j])
					} else if j == len(winnerName)-2 {
						fmt.Print(winnerName[j], " and ")
					} else {
						fmt.Print(winnerName[j], " , ")
					}

				}

			} else {
				fmt.Println("Please Play")
			}
			fmt.Println("and Score is ", winnerScore)

		}
	}
}
