package config

const filePath = "../data/apartments.csv"

func (c *Config) FilePath() string {
	if c.options.FilePath != "" {
		return c.options.FilePath
	}

	return filePath
}
