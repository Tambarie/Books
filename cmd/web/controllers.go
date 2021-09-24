package main

import "net/http"



func homepage(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("HomePage"))
}