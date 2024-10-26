package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
	// Render the course html page
	http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "static/contact.html")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Healthy"))
}

func main() {
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/courses", coursePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)
	http.HandleFunc("/health", healthCheck)

	// Start server and log the port being listened on
	log.Println("Starting server on :8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}