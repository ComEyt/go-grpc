package defalutSelectParam

const (
	defaultAddr = "127.0.0.1"
	defaultPort = 8000
)

type Server struct {
	Addr string
	Port int
}

// ServerOptions 用来生成默认参数
type ServerOptions struct {
	Addr string
	Port int
}

// NewServerOptions 创建默认的ServerOptions对象
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Addr: defaultAddr,
		Port: defaultPort,
	}
}

// NewServerWithOptions 通过ServerOptions创建一个Server对象
func NewServerWithOptions(opts *ServerOptions) *Server {
	return &Server{
		Addr: opts.Addr,
		Port: opts.Port,
	}
}
