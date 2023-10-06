package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//create a web server

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		d, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %s", d)
	})

	http.HandleFunc("/goodby", func(http.ResponseWriter, *http.Request) {
		log.Println("goodby")
	})

	http.ListenAndServe(":9090", nil)

}
