package dayOne

import (
	file "adventOfCode2022/files"
	"fmt"
	"strconv"
)

func countCalories(elfBackpacks map[int][]int) map[int]int {
	sumOfCaloriesPerElf := make(map[int]int)
	for i, elfBackpack := range elfBackpacks {
		caloriesSummarized := 0
		for _, calories := range elfBackpack {
			caloriesSummarized += calories
		}
		sumOfCaloriesPerElf[i] = caloriesSummarized
	}
	return sumOfCaloriesPerElf
}

func GetElfWithMostCalories(filepath string) {
	input := file.ReadFromFile(filepath)
	elfBackpacks := splitInputbyElves(input)
	fmt.Println(getMax(elfBackpacks))

}

func getMax(elfBackpacks map[int][]int) int {
	summarizedBackPacks := countCalories(elfBackpacks)
	max := 0
	for _, summarizedBackpack := range summarizedBackPacks {
		if summarizedBackpack > max {
			max = summarizedBackpack
		}
	}
	return max
}

func splitInputbyElves(input []string) map[int][]int {
	elves := make(map[int][]int)
	elfCounter := 1
	for _, inputLine := range input {
		if inputLine == "" {
			elfCounter = elfCounter + 1
		} else {
			inputInteger, err := strconv.Atoi(inputLine)
			if err == nil {
				elves[elfCounter] = append(elves[elfCounter], inputInteger)
			}
		}
	}
	return elves
}
