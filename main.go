package main

import (
	"log"
	"net/http"
)

func main(){
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello, World!"))
	})


	log.Print("Listen at port 4000")
	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}