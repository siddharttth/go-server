package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		phone := r.FormValue("phone")
		address := r.FormValue("address")
		message := r.FormValue("message")
		fmt.Fprintf(w, "Name: %s\nEmail: %s\nPhone: %s\nAddress: %s\nMessage: %s", name, email, phone, address, message)
	} else {
		http.ServeFile(w, r, "static/form.html")
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/form", formHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Println("Server is running on port 3000...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
