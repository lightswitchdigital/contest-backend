package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"parasource/contest/src/server"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	config := server.ServerConfig{
		Host: host,
		Port: port,
	}
	s, err := server.NewServer(config)
	if err != nil {
		logrus.Fatalf("error creating server: %v", err)
	}

	stopCh := make(chan struct{}, 1)

	s.Run(stopCh)
}
