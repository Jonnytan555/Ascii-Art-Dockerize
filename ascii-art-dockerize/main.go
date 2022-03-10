package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type AsciiArt struct{}

type AsciiArtPost struct {
	Ascii_Art_Text string
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/", GetHandler)
	http.HandleFunc("/ascii-art", PostHandler)

	fmt.Printf("Starting server at port 3000\n")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {

	p := strings.Split(r.URL.Path, "/")[1:]

	if len(p) == 1 && p[0] == "" {
		tpl := template.Must(template.ParseFiles("templates/form.html"))
		//tpl, _ := template.ParseGlob("*/*.html")

		tpl.ExecuteTemplate(w, "index.html", "")
	} else {
		http.Error(w, "404 Status Not Found", http.StatusNotFound)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	//tpl, _ := template.ParseGlob("*/*.html")
	tpl := template.Must(template.ParseFiles("templates/home.html"))

	tpl.ExecuteTemplate(w, "form.html", nil)

}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	tpl := template.Must(template.ParseFiles("templates/form.html"))
	//tpl, _ := template.ParseGlob("*/*.html")

	Input := r.FormValue("Input")
	Font := r.FormValue("Font")
	if len(Input) == 0 {
		http.Error(w, "400 Bad Request", http.StatusNotFound)
		return

	}

	Ascii_output, Open_Error := ToAscii(Input, Font)
	if Open_Error != nil {
		http.Error(w, "500 Internal Service Error", http.StatusInternalServerError)
		return
	}

	Art := string(Ascii_output)

	Ascii_Art_Text_struct := AsciiArtPost{Ascii_Art_Text: Art}

	tpl.ExecuteTemplate(w, "form.html", Ascii_Art_Text_struct)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../images/favicon.ico")
}

func ToAscii(input, font string) (string, error) {

	var Open_Error error

	output, err := os.Open(font)
	if err != nil {
		fmt.Println(err.Error())
		Open_Error = err
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

	return aString, Open_Error

}
