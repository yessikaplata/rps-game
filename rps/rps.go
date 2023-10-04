package rps

import (
	"strconv"
)

const (
	ROCK     = 0 //Piedra. Vence a las tijeras (tijeras + 1) % 3 = 0
	PAPER    = 1 //Papel. Vence a la piedra (piedra + 1) % 3 = 1
	SCISSORS = 2 //Piedra. Vence a las papel (papel + 1) % 3 = 2
)

var winMessages = []string{
	"You got it!",
	"Good Job!",
	"Well done!",
}

var loseMessages = []string{
	"Try again!",
	"You are out of luck!",
	"Is not your day!",
}

type Round struct {
	Message            string `json:"message"`
	ComputerChoice     int    `json:"computer_choice"`
	ComputerChoiceName string `json:"computer_choice_name"`
	RoundResult        string `json:"round_result"`
	ComputerScore      string `json:"computer_score"`
	PlayerScore        string `json:"player_score"`
}

type RandomGenerator interface {
	Intn(n int) int
}

type RPSGame struct {
	Random        RandomGenerator
	ComputerScore int
	PlayerScore   int
}

func NewRPSGame(random RandomGenerator) RPSGame {
	return RPSGame{
		Random:        random,
		ComputerScore: 0,
		PlayerScore:   0,
	}
}

func (rps *RPSGame) PlayRound(playerValue int) Round {
	computerValue := rps.Random.Intn(3)
	messageNumber := rps.Random.Intn(3)
	round := Round{}
	switch computerValue {
	case ROCK:
		round.ComputerChoice = ROCK
		round.ComputerChoiceName = "The Computer chosen Rock"
	case PAPER:
		round.ComputerChoice = PAPER
		round.ComputerChoiceName = "The Computer chosen Paper"
	case SCISSORS:
		round.ComputerChoice = SCISSORS
		round.ComputerChoiceName = "The Computer chosen Scissors"
	}
	if playerValue == computerValue {
		round.RoundResult = "It is a draw"
		round.Message = loseMessages[messageNumber]
	} else if playerValue == (computerValue+1)%3 {
		round.RoundResult = "You have won"
		round.Message = winMessages[messageNumber]
		rps.PlayerScore++
	} else {
		round.RoundResult = "You have lost"
		round.Message = loseMessages[messageNumber]
		rps.ComputerScore++
	}
	round.ComputerScore = strconv.Itoa(rps.ComputerScore)
	round.PlayerScore = strconv.Itoa(rps.PlayerScore)
	return round
}
