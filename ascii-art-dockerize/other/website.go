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

	// Additional: Get CSS working
	//fs := http.FileServer(http.Dir("css"))
	//http.Handle("/css/", http.StripPrefix("/css/", fs))
	//http.HandleFunc("/home", HomeHandler)

var tpl *template.Template

func toAscii(input, font string) string {
	output, err := os.Open(font)
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

	arguments := input
	var argRune []rune
	for _, srune := range arguments {
		if srune != 10 {
			argRune = append(argRune, srune)
		} else {
			argRune = append(argRune, '\\')
			argRune = append(argRune, 'n')
		}
	}
	arguments = string(argRune)
	// Check if user input contains a newline character
	// if so, split the arguments by this newline character
	argsChecked := strings.Split(arguments, "\\n")
	output_file := make([]string, 8)
	//var output_file []string
	// For each charcter of the arguements, loop and print out each line
	fmt.Println(len(argsChecked))
	for _, str := range argsChecked {
		fmt.Println(str)
	}
	for _, word := range argsChecked {
		for j := 0; j < 8; j++ {
			for i := range word {
				output_file = append(output_file, aMap[int(word[i])][j])
				fmt.Println(i, j, len(aMap[int(word[i])][j]))
				fmt.Println(len(aMap[int(word[i])]))
				fmt.Println(len(word))
			}
			fmt.Println()
		}
	}
	for _, str := range output_file {
		fmt.Println(str)
	}
	aString := strings.Join(output_file, "")
	fmt.Println(aString)
	//fmt.Println(aString)
	return aString

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	Input := r.FormValue("Input")
	Font := r.FormValue("Font")

	var Art string

	//if strings.Contains(Input, "\\n") {
	//	Art = toAscii(string(aSlice), Font)
	//} else {
	Art = toAscii(Input, Font)

	tpl.ExecuteTemplate(w, "form.html", Art)
	fmt.Fprintf(w, "%s", Art)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
		return
	}

	Art := ""
	tpl.ExecuteTemplate(w, "form.html", Art)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "there was an error parsing page")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func main() {

	tpl, _ = tpl.ParseGlob("templates/*.html")
	fileServer := http.FileServer(http.Dir("./static"))
	//myDir := http.Dir("/home/jonnytan555/ascii-art-web/images")
	//myHandler := http.FileServer(myDir)
	//http.Handle("/", myHandler)
	http.Handle("/", fileServer)
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/index", IndexHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/error", errorHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
