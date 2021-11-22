package phttp

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"ip-proxy/pkg/config"
	"ip-proxy/pkg/report"
	"ip-proxy/pkg/utils"
	"net"
	"net/textproto"
	"net/url"
	"strings"
)

type conn struct {
	rwc         net.Conn
	brc         *bufio.Reader
	server      *Server
	requestTime int64
}

// serve tunnel the client connection to remote host
func (c *conn) serve() {
	defer c.rwc.Close()

	rawHttpRequestHeader, remote, _, proxyClientIp, _, isHttps, err := c.getTunnelInfo()
	if err != nil {
		return
	}

	//	originUrl := c.rwc.RemoteAddr().String()

	var uid string

	//if isTunnel && proxyClientIp != "" {
	//	// 特殊处理，如果遇到隧道代理服务器，直接转发不进行授权等操作！
	//} else {
	//	uid = auth.HttpAuthorization(credential, originUrl)
	//
	//	if config.NoAuth == 0 {
	//		if uid == "" {
	//			c.rwc.Write([]byte("HTTP/1.1 407 Proxy Authentication Required\r\nProxy-Authenticate: Basic realm=\"*\"\r\n\r\n"))
	//			return
	//		}
	//	}
	//
	//	if uid != "" {
	//		limit := limit.CheckFlowControl(uid)
	//
	//		if limit {
	//			c.rwc.Write([]byte("HTTP/1.1 503 Service Unavailable\r\n"))
	//			return
	//		}
	//	}
	//}

	remoteConn, err := net.Dial("tcp", remote)
	if err != nil {
		return
	}

	if isHttps {
		// if https, should sent 200 to client
		_, err = c.rwc.Write([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
		if err != nil {
			return
		}
	} else {
		// if not https, should sent the request header to remote
		_, err = rawHttpRequestHeader.WriteTo(remoteConn)
		if err != nil {
			return
		}
	}

	if proxyClientIp == "" {
		proxyClientIp = c.rwc.RemoteAddr().String()
	}

	wRequest := config.WRequest{
		DeviceNum: config.Number,
		Uid:       uid,
		OriginIp:  proxyClientIp,
		ProxyIp:   config.Ip + ":" + fmt.Sprint(config.Port),
		RemoteIp:  remoteConn.RemoteAddr().String(),
		EventTime: c.requestTime,
	}

	report.W(wRequest)

	// build bidirectional-streams
	c.tunnel(remoteConn)
}

// getClientInfo parse client request header to get some information:
func (c *conn) getTunnelInfo() (rawReqHeader bytes.Buffer, host, credential, proxyClientIp string, isTunnel, isHttps bool, err error) {
	tp := textproto.NewReader(c.brc)

	// First line: GET /index.html HTTP/1.0
	var requestLine string
	if requestLine, err = tp.ReadLine(); err != nil {
		return
	}

	method, requestURI, _, ok := parseRequestLine(requestLine)
	if !ok {
		err = &BadRequestError{"malformed HTTP request"}
		return
	}

	// https request
	if method == "CONNECT" {
		isHttps = true
		requestURI = "http://" + requestURI
	}

	// get remote host
	uriInfo, err := url.ParseRequestURI(requestURI)
	if err != nil {
		return
	}

	// Subsequent lines: Key: value.
	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		return
	}

	credential = mimeHeader.Get(config.ProxyAuthorizationKey)
	tunnelCredential := mimeHeader.Get(config.TunnelXForKey)
	if tunnelCredential == config.TunnelXForValue {
		isTunnel = true
		proxyClientIp = mimeHeader.Get(config.XTunnelForwardedFor)
	}

	if uriInfo.Host == "" {
		host = mimeHeader.Get("Host")
	} else {
		if strings.Index(uriInfo.Host, ":") == -1 {
			host = uriInfo.Host + ":80"
		} else {
			host = uriInfo.Host
		}
	}

	// rebuild http request header
	rawReqHeader.WriteString(requestLine + "\r\n")
	for k, vs := range mimeHeader {
		for _, v := range vs {
			if utils.In(k, config.IgnoreHeaderMap) {
				continue
			}
			rawReqHeader.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
		}
	}
	rawReqHeader.WriteString("\r\n")
	return
}

// tunnel http message between client and server
func (c *conn) tunnel(remoteConn net.Conn) {
	//inAddr := c.rwc.RemoteAddr().String()
	//inLocalAddr := c.rwc.LocalAddr().String()
	//
	//outAddr := remoteConn.RemoteAddr().String()
	//outLocalAddr := remoteConn.LocalAddr().String()

	// log.Printf("conn %s - %s - %s - %s connected", inAddr, inLocalAddr, outLocalAddr, outAddr)

	go func() {
		c.brc.WriteTo(remoteConn)
		//if err != nil {
		//	log.Println(err)
		//}
		remoteConn.Close()
	}()

	io.Copy(c.rwc, remoteConn)
	//if err != nil {
	//	log.Println(err)
	//}

	// log.Printf("conn %s - %s - %s -%s released", inAddr, inLocalAddr, outLocalAddr, outAddr)
}

func parseRequestLine(line string) (method, requestURI, proto string, ok bool) {
	s1 := strings.Index(line, " ")
	s2 := strings.Index(line[s1+1:], " ")
	if s1 < 0 || s2 < 0 {
		return
	}
	s2 += s1 + 1
	return line[:s1], line[s1+1 : s2], line[s2+1:], true
}

type BadRequestError struct {
	what string
}

func (b *BadRequestError) Error() string {
	return b.what
}
