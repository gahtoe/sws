package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	fs := http.StripPrefix("/", http.FileServer(http.Dir("./static/")))

	http.Handle("/", fs)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        fs,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening on", s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
