package main

import (
	"testing"

	"./game"
	"./model"
)

func TestDicegame(t *testing.T) {
	var players []model.Players
	playerNames := []string{"Player1","Player2"}
	player1 := model.Players{Name: "Player1", Counter: 5, Score: 30}
	player2 := model.Players{Name: "Player2", Counter: 5, Score: 20}
	players = append(players, player1)
	players = append(players, player2)
	winnerName, _ := game.DiceGame(players)
	if winnerName[0] == playerNames[1] {
		t.Errorf("Expected winner name is Player1")
	}
	if winnerName[0] == "Player2" {
		t.Errorf("Expected winner name is Player1")
	}
	if game.NumberOfRounds != 5 {
		t.Errorf("Total number of rounds must be 5")

	}

}
