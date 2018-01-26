package main
import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)



	fmt.Println("starting	web	server	at	http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", router))
}