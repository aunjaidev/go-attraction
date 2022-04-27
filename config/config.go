package config

type Config struct {
	Server Server `config "server"`
}

type Server struct {
	Port        string `config "port"`
	ServiceName string `config "serviceName"`
	LogHttp     string `config "httpLog"`
}
