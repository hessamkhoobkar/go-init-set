package app

import "strconv"

func ValidateInput(input string) int {
	
	num, err := strconv.Atoi(input)
	
	if len([]rune(input)) != 1 {
		return 0
	}
	
	if err != nil {
		return 0
	}

	if num > 4 {
		return 0
	}

	return num
}