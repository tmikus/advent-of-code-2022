package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Move int

const (
	Rock     Move = 1
	Paper         = 2
	Scissors      = 3
)

type EncryptedMove int

const (
	X EncryptedMove = 1
	Y               = 2
	Z               = 3
)

type Round struct {
	Opponent           Move
	Me                 EncryptedMove
	ResultWhenRock     int
	ResultWhenPaper    int
	ResultWhenScissors int
}

func NewRound(opponent Move, me EncryptedMove) Round {
	return Round{
		opponent,
		me,
		-1,
		-1,
		-1,
	}
}

type MoveMapping struct {
	X Move
	Y Move
	Z Move
}

func evaluateDynamicStrategy(rounds []Round) int {
	points := 0
	for _, round := range rounds {
		points += evaluateDynamicRound(round)
	}
	return points
}

func evaluateDynamicRound(round Round) int {
	myMove := getMoveForRound(round)
	roundResult := getRoundResult(myMove, round.Opponent)
	return roundResult + int(myMove)
}

func getMoveForRound(round Round) Move {
	if round.Opponent == Rock {
		return getMoveForMe(round.Me, Scissors, Rock, Paper)
	}
	if round.Opponent == Paper {
		return getMoveForMe(round.Me, Rock, Paper, Scissors)
	}
	// Scissors
	return getMoveForMe(round.Me, Paper, Scissors, Rock)
}

func getMoveForMe(me EncryptedMove, lose, draw, win Move) Move {
	if me == X {
		return lose
	}
	if me == Y {
		return draw
	}
	return win
}

func evaluateFixedStrategy(rounds []Round) int {
	points := 0
	mapping := MoveMapping{Rock, Paper, Scissors}
	for _, round := range rounds {
		points += evaluateRound(round, mapping)
	}
	return points
}

func evaluateRound(round Round, mapping MoveMapping) int {
	myMove := mapEncryptedMove(round.Me, mapping)
	roundResult := getRoundResult(myMove, round.Opponent)
	return roundResult + int(myMove)
}

func getRoundResult(me, opponent Move) int {
	if me == Rock {
		if opponent == Scissors {
			return 6
		}
		if opponent == Rock {
			return 3
		}
		return 0
	}
	if me == Paper {
		if opponent == Rock {
			return 6
		}
		if opponent == Paper {
			return 3
		}
		return 0
	}
	// No need to check as we know that the only option is for me to have scissors
	// if me == Scissors {
	if opponent == Paper {
		return 6
	}
	if opponent == Scissors {
		return 3
	}
	return 0
}

func mapEncryptedMove(move EncryptedMove, mapping MoveMapping) Move {
	switch move {
	case X:
		return mapping.X
	case Y:
		return mapping.Y
	case Z:
		return mapping.Z
	}
	panic(fmt.Sprint("Unexpected encrypted move:", move))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rounds := make([]Round, 0)
	for scanner.Scan() {
		round := parseRoundFromLine(scanner.Text())
		rounds = append(rounds, round)
	}
	result := evaluateFixedStrategy(rounds)
	println("Part 1 result:", result)
	result = evaluateDynamicStrategy(rounds)
	println("Part 2 result:", result)
}

func parseRoundFromLine(line string) Round {
	lineParts := strings.Split(line, " ")
	return NewRound(
		parseMove(lineParts[0]),
		parseEncryptedMove(lineParts[1]),
	)
}

func parseEncryptedMove(part string) EncryptedMove {
	switch part {
	case "X":
		return X
	case "Y":
		return Y
	case "Z":
		return Z
	}
	panic(fmt.Sprint("Invalid encrypted move format: ", part))
}

func parseMove(part string) Move {
	switch part {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	}
	panic(fmt.Sprint("Invalid move format: ", part))
}
