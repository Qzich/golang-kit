package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ameteiko/golang-kit/log"
)

//
// Service class.
//
type Service struct {
	log              log.Logger
	router           HandlerProvider
	httpAddress      string
	httpReadTimeout  time.Duration
	httpWriteTimeout time.Duration
}

//
// NewService returns a new service instance.
//
func NewService(
	router HandlerProvider,
	log log.Logger,
	httpAddress string,
	httpReadTimeout time.Duration,
	httpWriteTimeout time.Duration,
) *Service {
	return &Service{
		router:           router,
		log:              log,
		httpAddress:      httpAddress,
		httpReadTimeout:  httpReadTimeout,
		httpWriteTimeout: httpWriteTimeout,
	}
}

//
// Run performs starts all application logic.
//
func (s *Service) Run() {
	s.log.Info(`Start listening address %v`, s.httpAddress)
	srv := http.Server{
		Addr:         s.httpAddress,
		Handler:      s.router.GetHTTPHandler(),
		ReadTimeout:  s.httpReadTimeout,
		WriteTimeout: s.httpWriteTimeout,
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-ch
		s.log.Info(`Graceful shutdown...`)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			s.log.Error(`%+v`, err)
		}

		s.log.Info(`Service has been stopped`)
	}()

	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.log.Error("%+v\n", err)
	}
}
