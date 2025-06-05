package config

import (
	"net/url"
	"strings"
)

func (c *Config) BaseUri(res string) string {
	u, err := url.Parse(c.SiteUrl())

	if err != nil {
		return res
	}

	return strings.TrimRight(u.EscapedPath(), "/") + res
}

func (c *Config) SiteUrl() string {
	if c.options.SiteUrl == "" {
		return "http://localhost:8080"
	}
	return c.options.SiteUrl
}

func (c *Config) ApiUri() string {
	return c.BaseUri(ApiUri)
}
