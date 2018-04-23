package main

import (
	"github.com/adriansr/grav/server"
)

func main() {
	s := server.New(":8080")
	s.Run(make(chan struct{}, 1))
}


