package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //参数解析
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello astaxie")
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("listenAndServer", err)
	}
}
