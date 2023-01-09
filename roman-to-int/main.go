package main

import (
	"log"
	"strings"
)

func main() {
	log.Println("Roman to Int")

	romanValue := "MCMXCIV"

	integerValue := romanToInt(romanValue)

	log.Println(romanValue, "->", integerValue)

}

func romanToInt(romanValue string) int {

	var romanIdentifier = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	romanValue = strings.Replace(romanValue, "CM", "CCCCCCCCC", -1)
	romanValue = strings.Replace(romanValue, "CD", "CCCC", -1)
	romanValue = strings.Replace(romanValue, "XC", "XXXXXXXXX", -1)
	romanValue = strings.Replace(romanValue, "XL", "XXXX", -1)
	romanValue = strings.Replace(romanValue, "IX", "IIIIIIIII", -1)
	romanValue = strings.Replace(romanValue, "IV", "IIII", -1)

	result := 0

	for _, roman := range romanValue {
		result += romanIdentifier[string(roman)]
	}

	return result
}
