package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, greeting("Code.education Rocks!"))
	})

	fmt.Println(http.ListenAndServe(":8000", nil))
}

func greeting(message string) string {
	return "<b>" + message + "</b>"
}
