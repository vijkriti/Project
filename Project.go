package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method id not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	// Print additional details
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

	// Additional details
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	gender := r.FormValue("gender")
	interests := r.FormValue("interests")
	age := r.FormValue("age")
	hobbies := r.FormValue("hobbies")
	feedback := r.FormValue("feedback")

	// Print additional details
	fmt.Fprintf(w, "Email = %s\n", email)
	fmt.Fprintf(w, "Phone = %s\n", phone)
	fmt.Fprintf(w, "Gender = %s\n", gender)
	fmt.Fprintf(w, "Interests = %s\n", interests)
	fmt.Fprintf(w, "Age = %s\n", age)
	fmt.Fprintf(w, "Hobbies = %s\n", hobbies)
	fmt.Fprintf(w, "Feedback = %s\n", feedback)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
