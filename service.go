package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/pat"
)

// fib returns a function that returns
// successive Fibonacci numbers.
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	string_number := r.URL.Query().Get(":number")
	number, _ := strconv.Atoi(string_number)
	fmt.Fprintln(w, fib(number))
}

func main() {
	r := pat.New()
	r.Get("/{number:[0-9]+}", IndexHandler)
	http.Handle("/", r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, nil)
}
