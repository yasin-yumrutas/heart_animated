package main

import "net/http"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "style.css")
	})

	http.HandleFunc("/index.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.js")
	})

	http.HandleFunc("/resim.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "resim.png")
	})

	http.ListenAndServe(":9998", nil)
}
