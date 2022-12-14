package api

import (
	"fmt"
	"go-interface/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config config.Config
	router *gin.Engine

	modules
}

func NewServer(config config.Config) (*Server, error) {
	gin.SetMode("debug")
	router := gin.New()
	router.Use(gin.Recovery())

	s := &Server{config: config, router: router}

	if err := s.initModules(s); err != nil {
		return nil, err
	}

	s.registerRoutes()

	return s, nil
}

func (s *Server) Run() error {
	if err := s.router.Run(fmt.Sprintf("0.0.0.0:%s", s.config.ListenAddr)); err != nil {
		return err
	}

	// server exit
	if err := s.fileSystem.Close(); err != nil {
		return err
	}

	return nil
}
