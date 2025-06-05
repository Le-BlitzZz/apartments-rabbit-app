package main

import (
	"context"
	"os"
	"sync"
	"syscall"

	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/config"
	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/presenter"
	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/uploader"
	log "github.com/sirupsen/logrus"

	"os/signal"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "Dataserver",
		Action: run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Panic(err)
	}
}

func run(ctx *cli.Context) error {
	conf, err := config.NewConfig(ctx)
	if err != nil {
		return err
	}
	defer conf.Shutdown()

	cctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	conf.InitDb()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		uploader.Start(cctx, conf)
	}()

	go func() {
		defer wg.Done()
		presenter.Start(cctx, conf)
	}()

	wg.Wait()

	return nil
}
