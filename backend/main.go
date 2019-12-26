package main

import (
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/"
)

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8080", nil)
}
 
func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}

"""
func Back_response(,r *http.Request) {
    return load_data()
}

func load_data() {
    return json_data
}
"""