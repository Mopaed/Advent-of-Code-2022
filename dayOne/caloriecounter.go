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

func GetSolutionsDayOne(filepath string) {
	input := file.ReadFromFile(filepath)
	elfBackpacks := splitInputbyElves(input)
	summarizedBackPacks := countCalories(elfBackpacks)

	fmt.Printf("Most Calories: %v \n", getMostCalories(summarizedBackPacks))
	fmt.Printf("Top 3 Calories summarized: %v\n", getSumOfCalorieesOfTopThreeElves(summarizedBackPacks))
}

func getMostCalories(summarizedBackPacks map[int]int) int {
	max := 0
	for _, summarizedBackpack := range summarizedBackPacks {
		if summarizedBackpack > max {
			max = summarizedBackpack
		}
	}
	return max

}

func getSumOfCalorieesOfTopThreeElves(elfBackpacks map[int]int) int {
	one := 0
	two := 0
	three := 0

	for _, calories := range elfBackpacks {
		if calories > one {
			three = two
			two = one
			one = calories
		} else if calories > two {
			three = two
			two = calories
		} else if calories > three {
			three = calories
		}
	}
	return one + two + three
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
