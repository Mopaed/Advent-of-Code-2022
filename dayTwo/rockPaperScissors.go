package dayTwo

import (
	file "adventOfCode2022/files"
	"fmt"
	"strings"
)

func GetSolutionsDayTwo(filepath string) {
	input := file.ReadFromFile(filepath)
	oponnentShapes, selfShapes := splitInputByOpponentAndSelf(input)
	fmt.Printf("Opponent shapes : %v\n", oponnentShapes)
	fmt.Printf("Self shapes: %v\n", selfShapes)
	fmt.Println(input)
}

func calculateOutcome(shapeOpponent string, shapeSelf string) int {

	return 0
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
