package main

import (
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/refresh", Refresh)

	//nil because not using any router framework
	log.Fatal(http.ListenAndServe(":3000", nil))
}
