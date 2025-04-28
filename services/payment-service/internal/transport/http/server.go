package http

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	wg     *sync.WaitGroup
	ctx    context.Context
	server *http.Server
	router *gin.Engine
}

func New(ctx context.Context, wg *sync.WaitGroup, cfg *Config, publicHandlers []Handler) *Server {
	s := &Server{
		wg: wg, ctx: ctx,
		server: &http.Server{
			Addr:              fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler:           nil,
			ReadHeaderTimeout: 10 * time.Second,
			ReadTimeout:       cfg.ReadTimeout,
			WriteTimeout:      cfg.WriteTimeout,
			IdleTimeout:       30 * time.Second,
		},
		router: gin.New(),
	}
	api := s.router.Group("")
	s.registerPublicHandlers(api, publicHandlers...)

	return s
}

func (s *Server) registerPublicHandlers(api *gin.RouterGroup, handlers ...Handler) {
	for _, h := range handlers {
		h.Register(api)
	}

	s.server.Handler = s.router
}

func (s *Server) Run() {
	s.wg.Add(1)
	zap.S().Infof("server listining: %s", s.server.Addr)

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.S().Error(err.Error())
	}
}

func (s *Server) Shutdown() error {
	zap.S().Info("Shutdown server...")
	zap.S().Info("Stopping http server...")

	ctx, cancel := context.WithTimeout(s.ctx, 30*time.Second)
	defer func() {
		cancel()
		s.wg.Done()
	}()

	if err := s.server.Shutdown(ctx); err != nil {
		zap.S().Fatal("Server forced to shutdown:", zap.Error(err))

		return err
	}

	zap.S().Info("Server successfully stopped.")

	return nil
}
