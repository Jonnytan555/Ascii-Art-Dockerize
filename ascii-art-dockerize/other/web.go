package web

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var tpl *template.Template

func toAscii(input, font string) string {
	output, err := os.Open(font)
	if err != nil {
		log.Fatalf("failed to open")
	}
	fmt.Println("1", input, "1.1", font)
	NewData := bufio.NewScanner(io.Reader(output))
	fmt.Println("2", NewData)
	// Scan the lines from the text file
	var lines []string

	NewData.Split(bufio.ScanLines)
	for NewData.Scan() {
		lines = append(lines, NewData.Text())
	}

	// Make a map for ascii characters --> lines
	aMap := make(map[int][]string)
	fmt.Println("3", aMap)
	// Fill up the map with the linesfmt.Println("3", aMap)
	// Each ascii character chnages when there is a line break
	i := 31

	for _, line := range lines {
		if line == "" {
			i++
		} else {
			aMap[i] = append(aMap[i], line)
		}
	}

	arguments := input

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

	aString := strings.Join(output_file, "")

	return aString
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	Input := r.FormValue("Input")
	Font := r.FormValue("Font")

	Art := toAscii(Input, Font)

	fmt.Println(Art)

	tpl.ExecuteTemplate(w, "form.html", Art)
	fmt.Fprintf(w, "%s", Art)

}

func main() {
	tpl, _ = tpl.ParseGlob("templates/*.html")
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
