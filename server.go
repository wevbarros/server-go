package main

import "net/http"

func main() {

	http.HandleFunc("/", Home)
	http.ListenAndServe(":3000", nil)

}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CR7 Goat"))

}
