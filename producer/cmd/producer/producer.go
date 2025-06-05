package main

import (
	"os"

	"github.com/Le-BlitzZz/apartments-rabbit-app/producer/internal/config"
	"github.com/Le-BlitzZz/apartments-rabbit-app/producer/internal/producer"
	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "Producer",
		Action: run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file-path",
				Value: "data/apartments.csv",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Panic(err)
	}
}

func run(c *cli.Context) error {
	conf, err := config.NewConfig(c)
	if err != nil {
		return err
	}
	defer conf.Shutdown()

	producer.Run(conf)

	return nil
}
