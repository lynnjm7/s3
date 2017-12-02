package main

import (
	"log"
	"net/http"
)

func Serve(options ArgumentOptions) {
	fs := http.FileServer(http.Dir(options.directory))
	http.Handle("/", fs)

	log.Println("Server started at localhost:" + options.port)
	http.ListenAndServe(":" + options.port, nil)
}

