package main

import (
	"log"
	"net/http"
	"time"

	".."
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/download", vault.DownloadTruststore)
	http.Handle("/", r)
	srv := &http.Server{
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
