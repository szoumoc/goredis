package main

import (
	"fmt"
	"log/slog"
	"net"
)

const DefaultServeraddr = ":8080"

type Config struct {
	Host             string
	ListenServeraddr string
}

type Server struct {
	Config
	peers     map[Peer]bool
	ln        net.Listener
	addPeerCh chan *Peer
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

	go s.loop()

	return s.acceptLoop()
}

func (s *Server) loop() error {
	for {
		select {
		case peer := <-s.addPeerCh:
			s.peers[*peer] = true
		default:
			fmt.Println("foo")
		}
	}
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
	peer := NewPeer(conn)
	s.addPeerCh <- peer
	go peer.readLoop()
}

func main() {
	// This is a placeholder for the main function.

}
