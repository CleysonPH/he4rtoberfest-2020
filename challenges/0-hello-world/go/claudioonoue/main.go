package main

import "fmt"

func main() {
	textOuput := []string{
		"▓█▀▀▀▀▀▀▀▀▀█▓▌░▄▄▄▄▄░",
		"▓█░░▀░░▀▄░░█▓▌░█▄▄▄█░",
		"▓█░░▄░░▄▀░░█▓▌░█▄▄▄█░",
		"▓█▄▄▄▄▄▄▄▄▄█▓▌░█████░",
		"░░░▄▄███▄▄░░░░░█████░",
	}

	for _, line := range textOuput {
		fmt.Println(line)
	}

	fmt.Println("Hello He4rtoberfest!!!")
}
