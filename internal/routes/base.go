package routes

import (
	"net/http"
	"sandbox/internal/entities"
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
		if err := s.service.PrepareState(ctx.Context()); err != nil {
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
	s.httpEngine.Get("/api/messages", func(ctx *fiber.Ctx) error {
		var pg entities.Pagination
		if err := ctx.QueryParser(&pg); err != nil {
			return ctx.Status(http.StatusBadRequest).SendString(err.Error())
		}
		msgList, err := s.service.GetMessages(ctx.Context(), &pg)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		return ctx.JSON(msgList)
	})
	s.httpEngine.Post("/api/message/:id", func(ctx *fiber.Ctx) error {
		pID, err := ctx.ParamsInt("id")
		if err != nil {
			return ctx.Status(http.StatusBadRequest).SendString("invalid message id")
		}
		var msg entities.EditChatMessage
		if err = ctx.BodyParser(&msg); err != nil {
			return ctx.Status(http.StatusBadRequest).SendString(err.Error())
		}
		if err = s.service.UpdateMessageByID(ctx.Context(), uint64(pID), &msg); err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		return ctx.SendString("ok")
	})
	s.httpEngine.Get("/api/message/:id", func(ctx *fiber.Ctx) error {
		pID, err := ctx.ParamsInt("id")
		if err != nil {
			return ctx.Status(http.StatusBadRequest).SendString("invalid message id")
		}
		msg, err := s.service.GetMessageByID(ctx.Context(), uint64(pID))
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		return ctx.JSON(msg)
	})
	s.httpEngine.Delete("/api/message/:id", func(ctx *fiber.Ctx) error {
		pID, err := ctx.ParamsInt("id")
		if err != nil {
			return ctx.Status(http.StatusBadRequest).SendString("invalid message id")
		}
		if err = s.service.DeleteMessageID(ctx.Context(), uint64(pID)); err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		return ctx.SendString("ok")
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
