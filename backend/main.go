package main

import "net/http"

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
