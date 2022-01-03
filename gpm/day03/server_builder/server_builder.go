package server_builder

import (
	"crypto/tls"
	"errors"
	"fmt"
	"time"
)

// 需要新建一个服务器配置，服务器和端口是必须参数
// 协议、时间、最大连接数、TLS 是可选参数

// Config 是每个 Server 独有的
type Config struct {
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

// 地址与端口是所有服务共有的
type Server struct {
	Addr   string
	Port   int
	Config *Config
}

func TryNew(addr string, port int, conf *Config) (*Server, error) {
	if len(addr) == 0 {
		return nil, errors.New("create new server failed, no addr")
	}

	if port < 0 || port > 65536 {
		return nil, errors.New("create new server failed, error port")
	}

	ns := Server{Port: port, Addr: addr, Config: conf}
	return &ns, nil
}

type Server2 struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

type ServerBuilder struct {
	Server2
}

func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
	sb.Server2.Addr = addr
	sb.Server2.Port = port
	// 其它代码设置其它成员的默认值
	return sb
}
func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
	sb.Server2.Protocol = protocol
	return sb
}
func (sb *ServerBuilder) WithMaxConn(maxconn int) *ServerBuilder {
	sb.Server2.MaxConns = maxconn
	return sb
}
func (sb *ServerBuilder) WithTimeOut(timeout time.Duration) *ServerBuilder {
	sb.Server2.Timeout = timeout
	return sb
}
func (sb *ServerBuilder) WithTLS(tls *tls.Config) *ServerBuilder {
	sb.Server2.TLS = tls
	return sb
}
func (sb *ServerBuilder) Build() Server2 {
	return sb.Server2
}

func TryBuilder() {
	sb := ServerBuilder{}
	server := sb.Create("127.0.0.1", 8080).
		WithProtocol("udp").
		WithMaxConn(1024).
		WithTimeOut(30 * time.Second).
		Build()
	fmt.Printf("%v\n", server)
}
