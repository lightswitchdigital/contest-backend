package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net"
)

type ServerConfig struct {
	Host string
	Port string
}

type Server struct {
	config  ServerConfig
	handler *Handler
	gin     *gin.Engine
}

func NewServer(conf ServerConfig) (*Server, error) {
	r := gin.Default()

	s := &Server{
		config:  conf,
		handler: NewHandler(),
	}
	s.setGinHandlers(r)

	s.gin = r

	return s, nil
}

func (s *Server) setGinHandlers(r *gin.Engine) {
	r.GET("/", s.handler.HandleIndex)
}

func (s *Server) Run(stopCh chan struct{}) {
	address := net.JoinHostPort(s.config.Host, s.config.Port)

	go func() {
		err := s.gin.Run(address)
		if err != nil {
			logrus.Fatalf("error running gin server: %v", err)
		}
	}()

	for {
		select {
		case <-stopCh:
			return
		}
	}
}
