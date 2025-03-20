package routes

import (
	"net/http"
	"sandbox/internal/entities"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) getAllMessages(ctx *fiber.Ctx) error {
	msgList, err := s.service.GetAllMessages(ctx.Context())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.JSON(msgList)
}

func (s *Server) queryMessages(ctx *fiber.Ctx) error {
	var pg entities.Pagination
	if err := ctx.QueryParser(&pg); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}
	msgList, err := s.service.GetMessages(ctx.Context(), &pg)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.JSON(msgList)
}

func (s *Server) updateMessageByID(ctx *fiber.Ctx) error {
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
}

func (s *Server) getMessageByID(ctx *fiber.Ctx) error {
	pID, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("invalid message id")
	}
	msg, err := s.service.GetMessageByID(ctx.Context(), uint64(pID))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.JSON(msg)
}

func (s *Server) deleteMessageByID(ctx *fiber.Ctx) error {
	pID, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("invalid message id")
	}
	if err = s.service.DeleteMessageID(ctx.Context(), uint64(pID)); err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.SendString("ok")
}
