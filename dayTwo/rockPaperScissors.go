package dayTwo

import (
	file "adventOfCode2022/files"
	"fmt"
	"strings"
)

const (
	win      int = 6
	draw     int = 3
	loss     int = 0
	rock     int = 1
	paper    int = 2
	scissors int = 3
)

func GetSolutionsDayTwo(filepath string) {
	input := file.ReadFromFile(filepath)
	opponentShapes, selfShapes := splitInputByOpponentAndSelf(input)
	fmt.Printf("Total Score by exact strategy guide: %v\n", calculateOutcome(opponentShapes, selfShapes))
	fmt.Printf("Total Score by changed strategy guide: %v\n", calculateOutcomeWithdifferentStrategy(opponentShapes, selfShapes))
}

func calculateOutcome(opponentShapes []string, selfShapes []string) int {
	if len(opponentShapes) != len(selfShapes) {
		return 0
	}
	outcome := 0
	for i := 0; i < len(opponentShapes); i++ {
		switch opponentShape := opponentShapes[i]; {
		case opponentShape == "A":
			switch selfShape := selfShapes[i]; {
			case selfShape == "X":
				outcome = outcome + draw + rock
			case selfShape == "Y":
				outcome = outcome + win + paper
			case selfShape == "Z":
				outcome = outcome + loss + scissors
			}
		case opponentShape == "B":
			switch selfShape := selfShapes[i]; {
			case selfShape == "X":
				outcome = outcome + loss + rock
			case selfShape == "Y":
				outcome = outcome + draw + paper
			case selfShape == "Z":
				outcome = outcome + win + scissors
			}
		case opponentShape == "C":
			switch selfShape := selfShapes[i]; {
			case selfShape == "X":
				outcome = outcome + win + rock
			case selfShape == "Y":
				outcome = outcome + loss + paper
			case selfShape == "Z":
				outcome = outcome + draw + scissors
			}
		}
	}
	return outcome
}

func calculateOutcomeWithdifferentStrategy(opponentShapes []string, strategyHints []string) int {
	// X = lose, Y = draw, Z = win
	if len(opponentShapes) != len(strategyHints) {
		return 0
	}
	outcome := 0
	for i := 0; i < len(opponentShapes); i++ {
		switch opponentShape := opponentShapes[i]; {
		case opponentShape == "A":
			switch selfShape := strategyHints[i]; {
			case selfShape == "X":
				outcome = outcome + loss + scissors
			case selfShape == "Y":
				outcome = outcome + draw + rock
			case selfShape == "Z":
				outcome = outcome + win + paper
			}
		case opponentShape == "B":
			switch selfShape := strategyHints[i]; {
			case selfShape == "X":
				outcome = outcome + loss + rock
			case selfShape == "Y":
				outcome = outcome + draw + paper
			case selfShape == "Z":
				outcome = outcome + win + scissors
			}
		case opponentShape == "C":
			switch selfShape := strategyHints[i]; {
			case selfShape == "X":
				outcome = outcome + loss + paper
			case selfShape == "Y":
				outcome = outcome + draw + scissors
			case selfShape == "Z":
				outcome = outcome + win + rock
			}
		}
	}
	return outcome
}

func splitInputByOpponentAndSelf(input []string) ([]string, []string) {
	var opponentShapes []string
	var selfShapes []string
	for _, shapesThisRound := range input {
		shapes := strings.Split(shapesThisRound, " ")
		opponentShapes = append(opponentShapes, shapes[0])
		selfShapes = append(selfShapes, shapes[1])
	}
	return opponentShapes, selfShapes
}
