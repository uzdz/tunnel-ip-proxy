package phttp

import (
	"bufio"
	"net"
)

type Server struct {
}

// NewServer create a proxy server
func NewServer() *Server {
	return &Server{}
}

// newConn create a conn to serve client request
func (s *Server) Process(rwc net.Conn, reader *bufio.Reader, requestTime int64) {
	run := &conn{
		server:      s,
		rwc:         rwc,
		brc:         reader,
		requestTime: requestTime,
	}
	run.serve()
}
