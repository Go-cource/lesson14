package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Write([]byte("Succes POST request"))
	} else if r.Method == http.MethodGet {
		w.Write([]byte(`{"status": "Succes Get"}`))
	}

}

func main() {
	http.HandleFunc("/", IndexHandler)
	fmt.Println("Server starting...")
	http.ListenAndServe(":8080", nil)
}
