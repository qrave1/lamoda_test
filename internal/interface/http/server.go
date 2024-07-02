package http

import (
	"fmt"
	"time"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/qrave1/lamoda_test/internal/interface/http/gen"
)

type Server struct {
	port string
	fb   *fiber.App
}

func NewServer(port string) *Server {
	return &Server{port: port}
}

func (s *Server) Run(api *API, pathToSpec string) error {
	s.fb = fiber.New()

	s.fb.Use(swagger.New(swagger.Config{
		FilePath: pathToSpec,
		Title:    "Warehouse API",
	}))

	middlewares := []gen.StrictMiddlewareFunc{
		RequestIDMiddleware,
	}

	srv := gen.NewStrictHandler(api, middlewares)
	gen.RegisterHandlers(s.fb, srv)

	errCh := make(chan error)
	go func() {
		errCh <- s.fb.Listen(fmt.Sprintf(":%s", s.port))
		close(errCh)
	}()
	select {
	case <-time.After(100 * time.Millisecond):
		return nil
	case err := <-errCh:
		return fmt.Errorf("error start fiber server. %w", err)
	}
}

func (s *Server) Shutdown(timeout time.Duration) {
	_ = s.fb.ShutdownWithTimeout(timeout)
}
