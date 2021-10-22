package main

import (
	"github.com/sirupsen/logrus"
	"parasource/contest/src/server"
)

func main() {
	config := server.ServerConfig{
		Host: "127.0.0.1",
		Port: "8080",
	}
	s, err := server.NewServer(config)
	if err != nil {
		logrus.Fatalf("error creating server: %v", err)
	}

	stopCh := make(chan struct{}, 1)

	s.Run(stopCh)
}
