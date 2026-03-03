package main

import (
	"fmt"
	//"html/template"
	//"log"
	"net/http"
)

func main() {
	host := "localhost"
	port := 8080
	fmt.Printf("Server running on host %s port %d", host, port)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", home)

}
func home(w http.ResponseWriter, req *http.Request) {

}
