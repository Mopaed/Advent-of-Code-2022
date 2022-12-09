package dayFour

import (
	file "adventOfCode2022/files"
	"fmt"
	"strconv"
	"strings"
)

type Shift struct {
	shiftOne string
	shiftTwo string
}

func GetSolutionsDayFour(filepath string) {
	input := file.ReadFromFile(filepath)
	shiftPlan := getShiftPlan(input)
	fmt.Printf("Number of shifts containing the second one: %v \n", getNumberOfSingleShiftsContainingTheOther(shiftPlan))

}

func getNumberOfSingleShiftsContainingTheOther(shiftPlan []Shift) int {
	containingShiftCounter := 0
	for _, shift := range shiftPlan {
		shiftBeginOne, shiftEndOne := getStartAndEndTime(shift.shiftOne)
		shiftBeginTwo, shiftEndTwo := getStartAndEndTime(shift.shiftTwo)
		if (shiftBeginOne <= shiftBeginTwo && shiftEndOne >= shiftEndTwo) || (shiftBeginOne >= shiftBeginTwo && shiftEndOne <= shiftEndTwo) {
			containingShiftCounter++
		}
	}
	return containingShiftCounter
}

func getNumberOfShiftsContainingAnother(shiftplan []Shift) int {
	return 0
}

func getStartAndEndTime(shift string) (int, int) {
	shiftBoundaries := strings.Split(shift, "-")
	begin, _ := strconv.Atoi(shiftBoundaries[0])
	end, _ := strconv.Atoi(shiftBoundaries[1])
	return begin, end
}

func getShiftPlan(input []string) []Shift {
	var shiftPlan []Shift
	for _, inputLine := range input {
		inputLineDivided := strings.Split(inputLine, ",")
		var shift Shift
		shift.shiftOne = inputLineDivided[0]
		shift.shiftTwo = inputLineDivided[1]
		shiftPlan = append(shiftPlan, shift)
	}
	return shiftPlan
}
