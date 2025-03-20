package sampler

import (
	"context"
	"fmt"
	"sandbox/internal/entities"
	"sandbox/internal/logger"
	"sandbox/internal/repository"
	"strings"
)

type Service struct {
	log        logger.AppLogger
	repo       *repository.Repo
	remoteHost string
}

func InitService(log logger.AppLogger, repo *repository.Repo, remoteHost string) *Service {
	return &Service{
		remoteHost: remoteHost,
		repo:       repo,
		log:        log.With(logger.WithService("sampler")),
	}
}

func (s *Service) GetMessages(ctx context.Context, pagination *entities.Pagination) ([]*entities.ChatMessage, error) {
	if pagination == nil {
		return nil, fmt.Errorf("pagination is required")
	}
	if pagination.PerPage <= 0 {
		return nil, fmt.Errorf("per_page should be greater than 0")
	}
	if pagination.Page <= 0 {
		return nil, fmt.Errorf("page should be greater than 0")
	}
	if pagination.PerPage > 20 {
		return nil, fmt.Errorf("per_page should be less than or equal to 20")
	}
	return s.repo.GetChatMessages(ctx, pagination)
}

func (s *Service) GetAllMessages(ctx context.Context) ([]*entities.ChatMessage, error) {
	msgList, err := s.repo.GetAllChatMessages(ctx)
	if err != nil {
		return nil, err
	}
	for i := range msgList {
		msgList[i].Message = strings.Repeat("*", len(msgList[i].Message))
	}
	return msgList, nil
}

func (s *Service) GetMessageByID(ctx context.Context, messageID uint64) (*entities.ChatMessage, error) {
	return s.repo.GetChatMessageByID(ctx, messageID)
}

func (s *Service) DeleteMessageID(ctx context.Context, messageID uint64) error {
	return s.repo.DeleteChatMessageByID(ctx, messageID)
}

func (s *Service) UpdateMessageByID(ctx context.Context, messageID uint64, message *entities.EditChatMessage) error {
	return s.repo.UpdateChatMessageByID(ctx, messageID, message)
}
