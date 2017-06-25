package main

import (
	"fmt"
	"io"
    "html/template"
    "net/http"
)

var words []string

func sentence(w http.ResponseWriter, r *http.Request) {
	// If GET method load form template and add last word
	lastWord := "NO WORD YET"
	if r.Method == "GET" {
		if len(words) > 0 {
			lastWord = words[len(words) - 1]
		}
        t, _ := template.ParseFiles("form.gtpl")
        t.Execute(w, lastWord)
    // If POST method, parse form and display
    } else {
        r.ParseForm()
        newWord := r.Form["word"][0]
        
        // Test if new word is already in the list
        inList := isWordInList(newWord, words)

        // Add the word to the list
        words = append(words, newWord)

        // Display result
        if inList {
        	io.WriteString(w, "YOU FINISHED THE SENTENCE : " + aggregateWords(words))
        } else {
        	io.WriteString(w, "THANK YOU, YOU ADDED : " + newWord)
        }
    }
}

func isWordInList(w string, words []string) bool {
	for _, v := range words {
		if v == w {
			return true
		}
	}
	return false
}

func aggregateWords(words []string) string {
	sentence := ""
	for _, v := range words {
		sentence += v + " "
	}
	fmt.Println(sentence)
	return sentence[:len(sentence)-1]
}

func main() {
	http.HandleFunc("/", sentence)
	http.ListenAndServe(":8000", nil)
}