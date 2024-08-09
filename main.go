package main

import (
	"fmt"
	_ "fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func pageHandler(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	pageName := params["page"]
	fileName := "static/html/"+ pageName+".html"
	fmt.Println(fileName)
	if  _, err := os.Stat(fileName); err != nil{
		fileName ="static/html/404.html"
	}
	http.ServeFile(w, r, fileName)

}

func main() {
	r := mux.NewRouter()
	fmt.Println("Listening")
	r.HandleFunc("/{page}", pageHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)

}