package main

import (
	"flag"
)

type ArgumentOptions struct {
	directory string 
	port string
}

func ParseArgs() ArgumentOptions {
	var argOptions ArgumentOptions
	flag.StringVar(&argOptions.directory, "dir", ".", "The directory to serve")
	flag.StringVar(&argOptions.port, "port", "8080", "Port for serving the directory")

	flag.Parse()

	return argOptions
}


