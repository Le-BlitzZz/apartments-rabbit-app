package config

const dataSize = 11986

func (c *Config) DataSize() int64 {
	if c.options.DataSize == 0 {
		return dataSize
	}
	return c.options.DataSize
}
