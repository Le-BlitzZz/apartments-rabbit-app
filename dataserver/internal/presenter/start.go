package presenter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Start(ctx context.Context, conf *config.Config) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()

	log.Info("presenter: started")

	router := gin.Default()

	API := router.Group(conf.ApiUri())
	registerRoutes(API)

	tcpSocket := fmt.Sprintf("%s:%d", conf.HttpHost(), conf.HttpPort())
	server := &http.Server{
		ReadHeaderTimeout: time.Minute,
		ReadTimeout:       -1,
		WriteTimeout:      -1,
		Handler:           router,
		Addr:              tcpSocket,
	}

	log.Infof("presenter: listening on %s", server.Addr)

	go StartHttp(server)

	<-ctx.Done()

	log.Info("presenter: shutting down")

	err := server.Close()
	if err != nil {
		log.Errorf("presenter: shutdown failed (%s)", err)
	}
}

func StartHttp(s *http.Server) {
	if err := s.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Info("presenter: shutdown complete")
		} else {
			log.Errorf("presenter: %s", err)
		}
	}
}
