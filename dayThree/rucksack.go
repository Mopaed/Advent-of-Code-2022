package dayThree

import (
	file "adventOfCode2022/files"
	"fmt"
	"strings"
	"unicode"
)

type Rucksack struct {
	bagOne string
	bagTwo string
}

func GetSolutionsDayThree(filepath string) {
	input := file.ReadFromFile(filepath)
	rucksacks := getRucksacks(input)
	fmt.Printf("Value of wrong packed goods: %v\n", getValueOfWrongPackedGoods(rucksacks))
}

func getValueOfWrongPackedGoods(rucksacks []Rucksack) int {
	valueOfWrongPackedGoods := 0
	for _, rucksack := range rucksacks {
		for i := 0; i < len(rucksack.bagOne); i++ {
			letters := strings.Split(rucksack.bagOne, "")
			if strings.Contains(rucksack.bagTwo, letters[i]) {
				valueOfWrongPackedGoods = valueOfWrongPackedGoods + getNumberValueOfLetter(letters[i])
				break
			}
		}
	}
	return valueOfWrongPackedGoods
}

func getRucksacks(input []string) []Rucksack {
	var rucksacks []Rucksack
	for _, backpack := range input {
		bagOne := backpack[0 : len(backpack)/2]
		bagTwo := backpack[len(backpack)-len(backpack)/2:]
		rucksack := Rucksack{
			bagOne: bagOne,
			bagTwo: bagTwo,
		}
		rucksacks = append(rucksacks, rucksack)
	}
	return rucksacks
}

func getNumberValueOfLetter(letter string) int {
	if len(letter) > 1 {
		return 0
	}
	letterRune := rune(letter[0])
	if unicode.IsUpper(letterRune) {
		return int(letterRune) - 38
	} else {
		return int(letterRune) - 96
	}
}
