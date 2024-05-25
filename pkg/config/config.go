package config

type CassiniConfig struct {
	Service ServiceConfig
}

type ServiceConfig struct {
	Name  string
	Port  int
	Url   string
	Tests []Test
}

type Test struct {
	Name   string
	Method string
	Status int
}
