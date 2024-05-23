package config

type CassiniConfig struct {
	Service ServiceConfig
}

type ServiceConfig struct {
	Name   string
	Port   int
	Url    string
	Method string
}
