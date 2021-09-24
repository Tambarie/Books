package main

import (
	"log"
	"net/http"
)





func main()  {
	
	mux := http.NewServeMux()

	server := http.Server{
		Addr: ":9007",
		Handler: mux,
	}

	mux.HandleFunc("/",homepage)



	log.Println("Starting server on port 9007")
	
	if err := server.ListenAndServe(); err != nil{
		log.Fatal(err)
	}
}