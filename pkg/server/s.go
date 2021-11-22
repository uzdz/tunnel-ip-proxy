package server

import (
	"bufio"
	phttp "ip-proxy/pkg/http"
	"ip-proxy/pkg/socks5"
	"log"
	"net"
	"strings"
	"time"
)

var P *S

type Server interface {
	Process(conn net.Conn, reader *bufio.Reader, requestTime int64)
}

func openS5() Server {
	authenticators := []socks5.Authenticator{&socks5.UserPassAuthenticator{}}
	conf := &socks5.Config{AuthMethods: authenticators}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	return server
}

func openHttp() Server {
	return phttp.NewServer()
}

type S struct {
	listener net.Listener
}

func (s *S) RunServer(port string) {

	s5Server := openS5()
	httpServer := openHttp()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
	}
	s.listener = listener

	log.Println("代理服务器启动成功...")

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				log.Println("正在关闭代理服务器...")
				break
			}
			log.Println(err)
			continue
		}

		requestTime := time.Now().Unix()
		bufConn := bufio.NewReader(conn)

		if v, err := bufConn.Peek(1); err == nil && 0 < v[0] && v[0] < 10 {
			go s5Server.Process(conn, bufConn, requestTime)
		} else {
			go httpServer.Process(conn, bufConn, requestTime)
		}
	}
}

func (s *S) Shutdown() bool {
	err := s.listener.Close()

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
