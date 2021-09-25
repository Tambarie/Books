package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	addr := flag.String("addr",":9007","Pass the network  address")
	flag.Parse()
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    *addr,
		Handler: mux,
	}

	staticFileserver := http.FileServer(http.Dir("../../ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static",staticFileserver))

	mux.HandleFunc("/", homepage)

	log.Printf("Starting server on port %s\n",*addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
