package socks5

import (
	"bufio"
	"fmt"
	"ip-proxy/pkg/config"
	"ip-proxy/pkg/report"
	"log"
	"net"
	"strconv"

	"golang.org/x/net/context"
)

const (
	socks5Version = uint8(5)
)

// Config is used to setup and configure a Server
type Config struct {
	// AuthMethods can be provided to implement custom authentication
	// By default, "auth-less" mode is enabled.
	// For password-based auth use UserPassAuthenticator.
	AuthMethods []Authenticator

	// Resolver can be provided to do custom name resolution.
	// Defaults to DNSResolver if not provided.
	Resolver NameResolver

	// Rules is provided to enable custom logic around permitting
	// various commands. If not provided, PermitAll is used.
	Rules RuleSet

	// Rewriter can be used to transparently rewrite addresses.
	// This is invoked before the RuleSet is invoked.
	// Defaults to NoRewrite.
	Rewriter AddressRewriter

	// BindIP is used for bind or udp associate
	BindIP net.IP

	// Logger can be used to provide a custom log target.
	// Defaults to stdout.
	Logger *log.Logger

	// Optional function for dialing out
	Dial func(ctx context.Context, network, addr string) (net.Conn, error)
}

// Server is reponsible for accepting connections and handling
// the details of the SOCKS5 protocol
type Server struct {
	config      *Config
	authMethods map[uint8]Authenticator
}

// New creates a new Server and potentially returns an error
func New(conf *Config) (*Server, error) {
	// Ensure we have at least one authentication method enabled
	//if len(conf.AuthMethods) == 0 {
	//	if conf.Credentials != nil {
	//		log.Println("用户密码授权")
	//		conf.AuthMethods = []Authenticator{&UserPassAuthenticator{conf.Credentials}}
	//	} else {
	//		log.Println("无授权模式")
	//		conf.AuthMethods = []Authenticator{&NoAuthAuthenticator{}}
	//	}
	//}

	// Ensure we have a DNS resolver
	if conf.Resolver == nil {
		conf.Resolver = DNSResolver{}
	}

	// Ensure we have a rule set
	if conf.Rules == nil {
		conf.Rules = PermitAll()
	}

	// Ensure we have a log target
	if conf.Logger == nil {
		conf.Logger = log.New(log.Writer(), "", log.LstdFlags)
	}

	server := &Server{
		config: conf,
	}

	server.authMethods = make(map[uint8]Authenticator)

	for _, a := range conf.AuthMethods {
		server.authMethods[a.GetCode()] = a
	}

	return server, nil
}

// ServeConn is used to serve a single connection.
func (s *Server) Process(conn net.Conn, bufConn *bufio.Reader, requestTime int64) {
	defer conn.Close()

	// Read the version byte
	version := []byte{0}
	if _, err := bufConn.Read(version); err != nil {
		return
	}

	// Ensure we are compatible
	if version[0] != socks5Version {
		return
	}

	// Authenticate the connection
	authContext, uid, err := s.authenticate(conn, bufConn, requestTime)
	if err != nil {
		return
	}

	request, err := NewRequest(bufConn)
	if err != nil {
		if err == unrecognizedAddrType {
			if err := sendReply(conn, addrTypeNotSupported, nil); err != nil {
				return
			}
		}
		return
	}
	request.AuthContext = authContext
	var clientIp string
	var clientPort int
	if client, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
		clientIp = client.IP.String()
		clientPort = client.Port
		request.RemoteAddr = &AddrSpec{IP: client.IP, Port: client.Port}
	}

	wRequest := config.WRequest{
		DeviceNum: config.Number,
		Uid:       uid,
		OriginIp:  clientIp + ":" + strconv.Itoa(clientPort),
		ProxyIp:   config.Ip + ":" + fmt.Sprint(config.Port),
		RemoteIp:  request.DestAddr.String(),
		EventTime: requestTime,
	}

	report.W(wRequest)

	// Process the client request
	if err := s.handleRequest(request, conn); err != nil {
		return
	}
}
