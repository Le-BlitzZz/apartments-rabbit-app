package config

const (
	httpHost = "0.0.0.0"
	httpPort = 8080
)

func (c *Config) HttpHost() string {
	if c.options.HttpHost == "" {
		return httpHost
	}
	return c.options.HttpHost
}

func (c *Config) HttpPort() int {
	if c.options.HttpPort == 0 {
		return httpPort
	}
	return c.options.HttpPort
}
