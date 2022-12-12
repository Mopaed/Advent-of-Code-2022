package main

import (
	"adventOfCode2022/dayFive"
	"adventOfCode2022/dayFour"
	"adventOfCode2022/dayOne"
	"adventOfCode2022/dayThree"
	"adventOfCode2022/dayTwo"
	"fmt"
)

func main() {
	fmt.Println("Start Day 1!")
	dayOne.GetSolutionsDayOne("resources/dayOne/calories.txt")
	fmt.Println("End Day 1!")
	fmt.Println("Start Day 2!")
	dayTwo.GetSolutionsDayTwo("resources/dayTwo/strategyGuide.txt")
	fmt.Println("End Day 2!")
	fmt.Println("Start Day 3!")
	dayThree.GetSolutionsDayThree("resources/dayThree/rucksacks.txt")
	fmt.Println("End Day 3!")
	fmt.Println("Start Day 4!")
	dayFour.GetSolutionsDayFour("resources/dayFour/shiftplan.txt")
	fmt.Println("End Day 4!")
	fmt.Println("Start Day 5!")
	dayFive.GetSolutionsDayFive("resources/dayFive/test.txt")
	fmt.Println("End Day 5!")
}
