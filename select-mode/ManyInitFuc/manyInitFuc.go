package ManyInitFuc

const (
	defaultAddr = "127.0.0.1"
	defaultPort = 8080
)

type Server struct {
	Addr string
	Port int
}

// NewServer 创建默认的Server对象
func NewServer() *Server {
	return &Server{
		Addr: defaultAddr,
		Port: defaultPort,
	}
}

// NewServerWithOptions 创建指定参数的Server对象
func NewServerWithOptions(addr string, port int) *Server {
	return &Server{
		Addr: addr,
		Port: port,
	}
}
