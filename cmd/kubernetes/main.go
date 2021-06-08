package main

import (
	"log"

	"github.com/melethron/kubernetes-service/internal/app/apiserver"
)

func main() {
	config := apiserver.NewConfig()
	//Create New apiserver
	s := apiserver.New(config)

	//Start apiserver
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
