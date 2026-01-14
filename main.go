package main

import (
	"log/slog"
	"net"
)

const DefaultServeraddr = ":8080"

type Config struct {
	Host             string
	ListenServeraddr string
	addPeerCh        chan *Peer
}

type Server struct {
	Config
	peers map[Peer]bool
	ln    net.Listener
}

func NewServer(cfg Config) *Server {
	if cfg.ListenServeraddr == "" {
		cfg.ListenServeraddr = DefaultServeraddr
	}
	return &Server{
		Config:    cfg,
		peers:     make(map[Peer]bool),
		addPeerCh: make(chan *Peer),
	}
}
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenServeraddr)
	if err != nil {
		return err
	}
	s.ln = ln
	return s.acceptLoop()
}

func (s *Server) acceptLoop() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			slog.Error("accept error", "error", err)
			continue
		}
		s.handleConnection(conn)
	}

}

func (s *Server) handleConnection(conn net.Conn) error {

}

func main() {
	// This is a placeholder for the main function.
}
