package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "There Will Be App")
	})
	log.Fatal(http.ListenAndServe(":8030", nil))
}
