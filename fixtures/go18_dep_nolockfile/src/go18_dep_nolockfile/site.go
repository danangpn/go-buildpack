package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ZiCog/shiny-thing/foo"
)

func main() {
	foo.Do()
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "go, world")
}
