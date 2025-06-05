package config

import (
	"reflect"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Options struct {
	BrokerServer      string
	BrokerUser        string
	BrokerPassword    string
	BrokerRawExchange string
	FilePath          string `flag:"file-path"`
}

func NewOptions(ctx *cli.Context) *Options {
	o := &Options{}

	if ctx == nil {
		return o
	}

	o.ApplyCliContext(ctx)

	return o
}

func (o *Options) ApplyCliContext(ctx *cli.Context) {
	v := reflect.ValueOf(o).Elem()

	for i := range v.NumField() {
		tagValue := v.Type().Field(i).Tag.Get("flag")

		if tagValue != "" {
			fieldValue := v.Field(i)

			switch t := fieldValue.Interface().(type) {
			case string:
				if ctx.IsSet(tagValue) || fieldValue.String() == "" {
					f := ctx.String(tagValue)
					fieldValue.SetString(f)
				}
			default:
				log.Println("Unsupported type:", t)
			}
		}
	}
}
