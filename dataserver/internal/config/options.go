package config

import "github.com/urfave/cli/v2"

type Options struct {
	BrokerServer              string
	BrokerUser                string
	BrokerPassword            string
	BrokerRawExchange         string
	BrokerTransformedExchange string
	BrokerRawQueue            string
	BrokerTransformedQueue    string
	DatabaseUser              string
	DatabasePassword          string
	DatabaseServer            string
	DatabaseName              string
	DatabaseTimeout           int
	DataSize                  int64
	HttpHost                  string
	HttpPort                  int
	SiteUrl                   string
}

func NewOptions(ctx *cli.Context) *Options {
	o := &Options{}

	if ctx == nil {
		return o
	}

	return o
}
