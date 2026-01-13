package main

import "net"

type Config struct {
	Host             string
	ListenServeraddr string
}

type Server struct {
	Config
	ln net.Listener
}

func NewServer(cfg Config) *Server {
	return &Server{
		Config: cfg,
	}
}

func main() {
	// This is a placeholder for the main function.
}
