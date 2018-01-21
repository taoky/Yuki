package main

import (
	"log"

	"github.com/knight42/Yuki/server"
)

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}
	s.Start()
}
