package routes

import (
	"fmt"
	"net/http"
	"sandbox/internal/config"
	"sandbox/internal/logger"
	"sandbox/internal/service/sampler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	conf       *config.AppConfig
	log        logger.AppLogger
	service    *sampler.Service
	httpEngine *fiber.App
}

// InitAppRouter initializes the HTTP Server.
func InitAppRouter(log logger.AppLogger, service *sampler.Service, conf *config.AppConfig) *Server {
	app := &Server{
		conf:       conf,
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
		if err := s.service.PrepareState(ctx.Context()); err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		return ctx.SendString("ok")
	})
	s.httpEngine.Get("/api/all_messages", s.getAllMessages)
	s.httpEngine.Get("/api/messages", s.queryMessages)
	s.httpEngine.Post("/api/message/:id", s.updateMessageByID)
	s.httpEngine.Get("/api/message/:id", s.getMessageByID)
	s.httpEngine.Delete("/api/message/:id", s.deleteMessageByID)
}

// Run starts the HTTP Server.
func (s *Server) Run() error {
	s.log.Info("starting HTTP server", logger.WithInt("port", s.conf.AppPort))
	return s.httpEngine.Listen(fmt.Sprintf(":%d", s.conf.AppPort))
}

func (s *Server) Stop() error {
	return s.httpEngine.Shutdown()
}
