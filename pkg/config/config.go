package config

type CassiniConfig struct {
	Server ServerConf
}

type ServerConf struct {
	Port   int
	Url    string
	Method string
}
