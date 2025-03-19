package routes

import (
	"net/http"
	"sandbox/internal/logger"
	"sandbox/internal/service/sampler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	appAddr    string
	log        logger.AppLogger
	service    *sampler.Service
	httpEngine *fiber.App
}

// InitAppRouter initializes the HTTP Server.
func InitAppRouter(log logger.AppLogger, service *sampler.Service, address string) *Server {
	app := &Server{
		appAddr:    address,
		httpEngine: fiber.New(fiber.Config{}),
		service:    service,
		log:        log.With(logger.WithService("http")),
	}
	app.httpEngine.Use(recover.New())
	app.initRoutes()
	return app
}

func (s *Server) initRoutes() {
	s.httpEngine.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})
	s.httpEngine.Get("/api/init", func(ctx *fiber.Ctx) error {
		if err := s.service.Init(ctx.Context()); err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		return ctx.SendString("ok")
	})
	s.httpEngine.Get("/api/all_messages", func(ctx *fiber.Ctx) error {
		msgList, err := s.service.GetAllMessages(ctx.Context())
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		return ctx.JSON(msgList)
	})
}

// Run starts the HTTP Server.
func (s *Server) Run() error {
	s.log.Info("starting HTTP server", logger.WithString("port", s.appAddr))
	return s.httpEngine.Listen(s.appAddr)
}

func (s *Server) Stop() error {
	return s.httpEngine.Shutdown()
}
