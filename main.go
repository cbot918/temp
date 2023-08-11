package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hihi\n"))
	})
	fmt.Println("listening: 6666")
	if err := http.ListenAndServe(":6666", nil); err != nil {
		log.Fatal(err)
	}
}
