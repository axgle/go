package main

import "net/http"

import "fmt"
import "html"
import "log"

func main() {
	//if you dont want to remember this method's  params,just run "godoc net/http" see more and copy it:)
	 http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8082", nil))
}