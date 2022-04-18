package config

type Web struct {
	port string
}

func (web Web) Port() string {
	return web.port
}
