package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>“Do not let the noise of others' opinions drown out your own inner voice. Have the courage to follow your heart and intuition; they somehow already know what you truly want to become - Steve Jobs”</h1>")
	fmt.Println("Someone hit me!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :80")
	http.ListenAndServe(":80", nil)
}
