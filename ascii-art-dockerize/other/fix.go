package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//input := "Hello World"
	font := "standard.txt"

	output, err := os.Open(font)
	if err != nil {
		fmt.Println(err.Error())
	}

	NewData := bufio.NewScanner(output)
	// Scan the lines from the text file
	var lines []string
	for NewData.Scan() {
		lines = append(lines, NewData.Text())
	}
	// Make a map for ascii characters --> lines
	aMap := make(map[int][]string)

	// Fill up the map with the lines
	// Each ascii character chnages when there is a line break
	i := 31
	for _, line := range lines {
		if line == "" {
			i++
		} else {
			aMap[i] = append(aMap[i], line)
		}
	}

	for _, line := range aMap[32] {
		fmt.Println(line)
	}

	/*
		arguments := input

		var box []byte

		for _, val := range arguments {
			if val != 10 && val != 13 {
				box = append(box, byte(val))
				continue
			}
			if val == 13 {
				continue
			}
			box = append(box, byte(92))
			box = append(box, byte(110))
		}

		// Check if user input contains a newline character
		// if so, split the arguments by this newline character

		argsChecked := strings.Split(string(box), "\\n")

		output_file := make([]string, 8)
		// For each charcter of the arguements, loop and print out each line
		for _, word := range argsChecked {
			for j := 0; j < 8; j++ {

				for i := range string(word) {

					output_file = append(output_file, aMap[int(word[i])][j])

				}
				output_file = append(output_file, "\n")
			}
		}

		aString := strings.Join(output_file, "")

		fmt.Println(aString)
	*/
}
