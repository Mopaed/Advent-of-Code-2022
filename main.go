package main

import (
	"adventOfCode2022/dayOne"
	"adventOfCode2022/dayTwo"
	"fmt"
)

func main() {
	fmt.Println("Start Day 1!")
	dayOne.GetSolutionsDayOne("resources/dayOne/calories.txt")
	fmt.Println("End Day 1!")
	fmt.Println("Start Day 2!")
	dayTwo.GetSolutionsDayTwo("resources/dayTwo/test.txt")
	fmt.Println("End Day 2!")
}
