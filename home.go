package main

import (
    "html/template"
    "net/http"
    "flag"
)

// Command-line flags.
var httpAddr = flag.String("http", ":8080", "Listen address")

var words []string

func home(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("index.gtpl")
    t.Execute(w, nil)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", home)
	http.ListenAndServe(*httpAddr, nil)
}