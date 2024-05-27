package config

type CassiniConfig struct {
	Service ServiceConfig
}

type ServiceConfig struct {
	Name           string
	Url            string
	IntervalUnit   string
	IntervalAmount int
	Tests          []Test
}

type Test struct {
	Name         string
	Method       string
	AssertStatus int
}
