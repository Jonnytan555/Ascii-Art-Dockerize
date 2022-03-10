func main() {

	if len(os.Args) != 4 || os.Args[3][:9] != "--output=" {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]")
		fmt.Println("EX: go run . something standard --output=<fileName.txt>")
	} else if os.Args[2] == "standard" || os.Args[2] == "shadow" || os.Args[2] == "thinkertoy" {
		getBanner(os.Args[2] + ".txt")
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]")
		fmt.Println("EX: go run . something standard --output=<fileName.txt>")
	}
}

func getBanner(banner string) {

	output, err := os.Open(banner)
	if err != nil {
		log.Fatalf("failed to open")
	}

	NewData := bufio.NewScanner(io.Reader(output))

	// Scan the lines from the text file
	var lines []string

	NewData.Split(bufio.ScanLines)
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

	arguments := os.Args[1]

	final := os.Args[3][9:]

	// Check if user input contains a newline character
	// if so, split the arguments by this newline character
	argsChecked := strings.Split(arguments, "\\n")

	output_file := make([]string, 8)
	// For each charcter of the arguements, loop and print out each line
	for _, word := range argsChecked {
		for j := 0; j < 8; j++ {
			for i := range word {
				output_file = append(output_file, aMap[int(word[i])][j])
			}
			output_file = append(output_file, "\n")
		}
	}

	// Rejoin the slice of strings andf convert to bytes
	aString := strings.Join(output_file, "")
	the_output := []byte(aString)

	// Create the output file and write the result to that file
	f, err := os.Create(final)
	if err != nil {
		fmt.Println(err)
	}

	w := bufio.NewWriter(f)
	w.WriteString(string(the_output))
	w.Flush()

}
