package main

import "fmt"
import "net/http"
import "html"

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar!")
}

func showUrl(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Hello, dari",html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo!")
	})
	http.HandleFunc("/index", index)
	http.HandleFunc("/url", showUrl)
	

	fmt.Println("starting	web	server	at	http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
