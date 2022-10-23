package main

import "fmt"

// go语言从语法上不支持默认参数的，所有为了实现既能通过默认参数创建对象。
// 又能够通过传递自定义参数创建对象，我们就需要通过一些编程技巧来实现。

// 选择模式
const (
	defaultAddr = "127.0.0.1"
	defaultPort = 8000
)

type Server struct {
	Addr string
	Port int
}

// ServerOptions 用来配置默认参数
type ServerOptions struct {
	Addr string
	Port int
}

type ServerOption interface {
	apply(*ServerOptions)
}

// FuncServerOption 定义这个结构体实现ServerOption接口是为了通过apply方法来为ServerOptions结构体单独配置某项参数
type FuncServerOption struct {
	f func(options *ServerOptions)
}

func (f FuncServerOption) apply(option *ServerOptions) {
	f.f(option)
}

func WithAddr(addr string) ServerOption {
	return FuncServerOption{
		f: func(options *ServerOptions) {
			options.Addr = addr
		}}
}
func WithPort(port int) ServerOption {
	return FuncServerOption{f: func(options *ServerOptions) {
		options.Port = port
	}}
}

func NewServer(opts ...ServerOption) *Server {
	options := ServerOptions{
		Addr: defaultAddr,
		Port: defaultPort,
	}
	for _, opt := range opts {
		opt.apply(&options)
	}
	return &Server{
		Addr: options.Addr,
		Port: options.Port,
	}
}

func main() {
	s2 := NewServer(WithAddr("localhost"), WithPort(8001))
	fmt.Println(s2)
}
