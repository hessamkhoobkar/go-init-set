package app

import (
	"fmt"
	"os"
)

func GetUserOption() string {

	fmt.Print("Please enter your choice: ")

	var input string

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "error"
	}

	return input
}